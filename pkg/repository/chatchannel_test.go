package repository_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/rand"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
)

var _ = Describe("ChatChannel Repository", func() {
	var pgRepo repository.ChatChannelRepository
	var id uuid.UUID
	BeforeEach(func(ctx SpecContext) {
		var err error

		pgRepo = repository.NewChatChannelPgxRepository(migrater)
		id, err = uuid.NewV7()
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create", func() {
		Context("no channel created", func() {
			It("should panic if the channel is nil", func(ctx SpecContext) {
				Expect(func() { pgRepo.Create(ctx, nil) }).To(Panic())
			})
			It("should return an error if the channel name is empty", func(ctx SpecContext) {
				createdChannel, outErr := pgRepo.Create(ctx, &chat.Channel{
					DimensionId: &id,
				})
				Expect(outErr).To(HaveOccurred())
				Expect(createdChannel).To(BeNil())
			})
		})
		Context("channel created", func() {
			var channel *chat.Channel
			var createdChannel *chat.Channel
			var outErr error
			BeforeEach(func(ctx SpecContext) {
				channel = &chat.Channel{
					Name:        faker.Username(),
					DimensionId: &id,
				}
			})
			It("should create a new chat channel", func(ctx SpecContext) {
				createdChannel, outErr = pgRepo.Create(ctx, channel)
			})

			It("should return an error if the channel already exists", func(ctx SpecContext) {
				createdChannel, outErr = pgRepo.Create(ctx, channel)
				createdChannel2, err := pgRepo.Create(ctx, channel)
				Expect(err).To(HaveOccurred())
				Expect(createdChannel2).To(BeNil())
			})

			It("should create a new chat channel if name and dimension combo existed but was deleted", func(ctx SpecContext) {
				createdChannel, outErr = pgRepo.Create(ctx, channel)
				err := pgRepo.Delete(ctx, &createdChannel.Id)
				Expect(err).NotTo(HaveOccurred())
				createdChannel, outErr = pgRepo.Create(ctx, channel)
			})

			AfterEach(func(ctx SpecContext) {
				Expect(outErr).NotTo(HaveOccurred())
				Expect(createdChannel).NotTo(BeNil())
				Expect(createdChannel.Id).NotTo(BeNil())
				Expect(createdChannel.Name).To(Equal(channel.Name))
				Expect(createdChannel.DimensionId).To(Equal(channel.DimensionId))
				Expect(createdChannel.CreatedAt).NotTo(BeNil())
				Expect(createdChannel.UpdatedAt).NotTo(BeNil())
				Expect(createdChannel.DeletedAt).To(BeNil())
			})
		})
	})

	Describe("GetById", func() {
		var channel *chat.Channel
		BeforeEach(func(ctx SpecContext) {
			var err error
			channel, err = pgRepo.Create(ctx, &chat.Channel{
				Name:        faker.Username(),
				DimensionId: &id,
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return an error if the channel ID is nil", func(ctx SpecContext) {
			outChannel, err := pgRepo.GetById(ctx, nil)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outChannel).To(BeNil())
		})
		It("should return the channel", func(ctx SpecContext) {
			outChannel, err := pgRepo.GetById(ctx, &channel.Id)
			Expect(err).NotTo(HaveOccurred())
			Expect(outChannel).NotTo(BeNil())
			Expect(outChannel.Id).To(Equal(channel.Id))
			Expect(outChannel.Name).To(Equal(channel.Name))
			Expect(outChannel.DimensionId).To(Equal(channel.DimensionId))
			Expect(outChannel.CreatedAt).To(Equal(channel.CreatedAt))
			Expect(outChannel.UpdatedAt).To(Equal(channel.UpdatedAt))
			Expect(outChannel.DeletedAt).To(BeNil())
		})
		It("should return an error if the channel does not exist", func(ctx SpecContext) {
			uuid, err := uuid.NewV7()
			Expect(err).NotTo(HaveOccurred())
			outChannel, err := pgRepo.GetById(ctx, &uuid)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outChannel).To(BeNil())
		})
	})

	Describe("GetAll", func() {
		It("should return all channels", func(ctx SpecContext) {
			count := rand.Intn(10) + 5
			for i := 0; i < count; i++ {
				id, err := uuid.NewV7()
				Expect(err).NotTo(HaveOccurred())
				channel, err := pgRepo.Create(ctx, &chat.Channel{
					Name:        faker.Username(),
					DimensionId: &id,
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(channel).NotTo(BeNil())

				if i%2 == 0 {
					err = pgRepo.Delete(ctx, &channel.Id)
					Expect(err).NotTo(HaveOccurred())
				}
			}

			channels, err := pgRepo.GetAll(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(channels).NotTo(BeNil())
			Expect(len(*channels) > (count / 2)).To(BeTrue())

			for _, channel := range *channels {
				Expect(channel.Id).NotTo(BeNil())
				Expect(channel.Name).NotTo(BeEmpty())
				Expect(channel.DeletedAt).To(BeNil())
			}
		})
	})

	Describe("Delete", func() {
		var channel *chat.Channel
		BeforeEach(func(ctx SpecContext) {
			var err error
			channel, err = pgRepo.Create(ctx, &chat.Channel{
				Name:        faker.Username(),
				DimensionId: &id,
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should delete the channel", func(ctx SpecContext) {
			err := pgRepo.Delete(ctx, &channel.Id)
			Expect(err).NotTo(HaveOccurred())

			outChannel, err := pgRepo.GetById(ctx, &channel.Id)
			Expect(err).To(Equal(pgx.ErrNoRows))
			Expect(outChannel).To(BeNil())
		})
		It("should err if the channel ID is nil", func(ctx SpecContext) {
			err := pgRepo.Delete(ctx, nil)
			Expect(err).To(Equal(repository.ErrDoesNotExist))
		})
		It("should return an error if the channel does not exist", func(ctx SpecContext) {
			uuid, err := uuid.NewV7()
			Expect(err).NotTo(HaveOccurred())
			err = pgRepo.Delete(ctx, &uuid)
			Expect(err).To(Equal(repository.ErrDoesNotExist))
		})
	})

	Describe("Save", func() {
		var channel *chat.Channel
		BeforeEach(func(ctx SpecContext) {
			var err error
			channel, err = pgRepo.Create(ctx, &chat.Channel{
				Name:        faker.Username(),
				DimensionId: &id,
			})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return an error if the channel does not exist", func(ctx SpecContext) {
			uuid, err := uuid.NewV7()
			Expect(err).NotTo(HaveOccurred())
			channel.Id = uuid
			_, err = pgRepo.Save(ctx, channel)
			Expect(err).To(Equal(repository.ErrDoesNotExist))
		})
		It("should update the channel", func(ctx SpecContext) {
			newId, err := uuid.NewV7()
			Expect(err).NotTo(HaveOccurred())

			channel.Name = faker.Username()
			channel.DimensionId = &newId
			updatedChannel, err := pgRepo.Save(ctx, channel)
			Expect(err).NotTo(HaveOccurred())
			Expect(updatedChannel).NotTo(BeNil())
			Expect(updatedChannel.Id).To(Equal(channel.Id))
			Expect(updatedChannel.Name).To(Equal(channel.Name))
			Expect(updatedChannel.DimensionId).To(Equal(channel.DimensionId))
			Expect(updatedChannel.CreatedAt).To(Equal(channel.CreatedAt))
			Expect(updatedChannel.UpdatedAt).NotTo(Equal(channel.UpdatedAt))
			Expect(updatedChannel.DeletedAt).To(BeNil())
		})
	})
})
