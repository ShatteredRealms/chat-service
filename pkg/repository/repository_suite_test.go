package repository_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	crepository "github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/ShatteredRealms/go-common-service/pkg/testsro"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus/hooks/test"
)

type initializeData struct {
	PgUrl string
}

var (
	hook *test.Hook

	pgCloseFunc func() error

	pgUrl    string
	migrater *crepository.PgxMigrater
)

func TestRepository(t *testing.T) {
	SynchronizedBeforeSuite(func() []byte {
		log.Logger, hook = test.NewNullLogger()

		var err error

		var pgPort string
		pgCloseFunc, pgPort, err = testsro.SetupPostgresWithDocker()
		Expect(err).NotTo(HaveOccurred())
		Expect(pgCloseFunc).NotTo(BeNil())

		cfg := config.DBConfig{
			ServerAddress: config.ServerAddress{
				Host: "localhost",
				Port: pgPort,
			},
			Name:     testsro.DbName,
			Username: testsro.Username,
			Password: testsro.Password,
		}

		data := initializeData{}
		data.PgUrl = cfg.PostgresDSN()

		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		Expect(enc.Encode(data)).To(Succeed())

		return buf.Bytes()
	}, func(ctx SpecContext, inBytes []byte) {
		log.Logger, hook = test.NewNullLogger()

		data := initializeData{}
		dec := gob.NewDecoder(bytes.NewBuffer(inBytes))
		Expect(dec.Decode(&data)).To(Succeed())

		pgUrl = data.PgUrl

		var err error
		migrater, err = crepository.NewPgxMigrater(ctx, pgUrl, "../../migrations")
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		log.Logger, hook = test.NewNullLogger()
	})

	SynchronizedAfterSuite(func() {
	}, func() {
		if pgCloseFunc != nil {
			pgCloseFunc()
		}
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}
