package repository_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
)

var _ = Describe("ChatLog", func() {
	var pgRepo repository.ChatLogRepository
	BeforeEach(func() {
		pgRepo = repository.NewChatLogPgxRepository(migrater)
	})
	Describe("AddMessage", func() {
		var msg *chat.Message
		BeforeEach(func() {
			msg = &chat.Message{
				SenderCharacterId: faker.UUIDHyphenated(),
				Content:           faker.Sentence(),
			}
		})
		It("should add a message to the chat log", func(ctx SpecContext) {
			id := uuid.New()
			id2 := uuid.New()
			err := pgRepo.AddMessage(ctx, &id, &id2, msg)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
