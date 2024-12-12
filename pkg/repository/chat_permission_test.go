package repository_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ShatteredRealms/chat-service/pkg/model/chat"
	"github.com/ShatteredRealms/chat-service/pkg/repository"
)

var _ = Describe("ChatPermission", func() {
	var pgRepo repository.ChatChannelPermissionRepository
	var channel *chat.Channel
	var characterId string
	BeforeEach(func(ctx SpecContext) {
		pgRepo = repository.NewChatChannelPermissionPgxRepository(migrater)

		ccRepo := repository.NewChatChannelPgxRepository(migrater)

		var err error
		channel, err = ccRepo.Create(ctx, &chat.Channel{
			Name:        faker.Username(),
			DimensionId: faker.UUIDHyphenated(),
		})
		Expect(err).NotTo(HaveOccurred())
		characterId = faker.UUIDHyphenated()
	})

	Describe("AddForCharacter", func() {
		It("should error if the character ID is empty", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, "", []*uuid.UUID{&channel.Id})
			Expect(err).To(HaveOccurred())
		})
		It("should error if the channel ID is nil", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{nil})
			Expect(err).To(HaveOccurred())
		})
		It("should do nothing if no channel IDs were given", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should add permissions for a character", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should err if permission already exists", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).NotTo(HaveOccurred())
			err = pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("RemForCharacter", func() {
		Context("no channel permissions removed", func() {
			It("should not error if the character ID is empty", func(ctx SpecContext) {
				err := pgRepo.RemForCharacter(ctx, "", []*uuid.UUID{&channel.Id})
				Expect(err).NotTo(HaveOccurred())
			})
			It("should not return an error if channel ID is nil", func(ctx SpecContext) {
				err := pgRepo.RemForCharacter(ctx, characterId, []*uuid.UUID{nil})
				Expect(err).NotTo(HaveOccurred())
			})
			It("should not return an error if no channels provided", func(ctx SpecContext) {
				err := pgRepo.RemForCharacter(ctx, characterId, []*uuid.UUID{})
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("channel permissions removed", func() {
			It("should remove permissions for a character", func(ctx SpecContext) {
				err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
				Expect(err).NotTo(HaveOccurred())

				err = pgRepo.RemForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetForCharacter", func() {
		It("should return nothing if the character ID is empty", func(ctx SpecContext) {
			channels, err := pgRepo.GetForCharacter(ctx, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(channels).NotTo(BeNil())
			Expect(len(*channels)).To(Equal(0))
		})
		It("should get permissions for a character with some", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).NotTo(HaveOccurred())

			channels, err := pgRepo.GetForCharacter(ctx, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(channels).NotTo(BeNil())
			Expect(len(*channels)).To(Equal(1))
		})
		It("should get permissions for a character with none", func(ctx SpecContext) {
			channels, err := pgRepo.GetForCharacter(ctx, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(channels).NotTo(BeNil())
			Expect(len(*channels)).To(Equal(0))
		})
	})

	Describe("GetAccessLevel", func() {
		It("should return false if the character ID is empty", func(ctx SpecContext) {
			hasAccess, err := pgRepo.GetAccessLevel(ctx, &channel.Id, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionNone))
		})
		It("should return false if the channel ID is nil", func(ctx SpecContext) {
			hasAccess, err := pgRepo.GetAccessLevel(ctx, nil, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionNone))
		})
		It("should return false if the character has no permissions", func(ctx SpecContext) {
			hasAccess, err := pgRepo.GetAccessLevel(ctx, &channel.Id, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionNone))
		})
		It("should return true if the character has permissions", func(ctx SpecContext) {
			err := pgRepo.AddForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).NotTo(HaveOccurred())

			hasAccess, err := pgRepo.GetAccessLevel(ctx, &channel.Id, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionReadSend))
		})
	})

	Describe("SaveForCharacter", func() {
		It("should return an error if the character ID is empty", func(ctx SpecContext) {
			err := pgRepo.SaveForCharacter(ctx, "", []*uuid.UUID{&channel.Id})
			Expect(err).To(HaveOccurred())
		})
		It("should return an error if the channel ID is nil", func(ctx SpecContext) {
			err := pgRepo.SaveForCharacter(ctx, characterId, []*uuid.UUID{nil})
			Expect(err).To(HaveOccurred())
		})
		It("should overwrite permissions for a character", func(ctx SpecContext) {
			err := pgRepo.SaveForCharacter(ctx, characterId, []*uuid.UUID{&channel.Id})
			Expect(err).NotTo(HaveOccurred())
			hasAccess, err := pgRepo.GetAccessLevel(ctx, &channel.Id, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionReadSend))

			err = pgRepo.SaveForCharacter(ctx, characterId, []*uuid.UUID{})
			Expect(err).NotTo(HaveOccurred())
			hasAccess, err = pgRepo.GetAccessLevel(ctx, &channel.Id, characterId)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasAccess).To(Equal(chat.PermissionNone))
		})
	})
})
