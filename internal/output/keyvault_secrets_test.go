package output_test

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/saliceti/yaq/internal/output"
	"github.com/saliceti/yaq/internal/pipeline"
	"github.com/saliceti/yaq/internal/testutils"
)

var _ = Describe("keyvault-secrets", Label("azure"), func() {
	var kvName string
	secretMap := pipeline.GenericMap{
		"yaq-keyvault-map-test-0": "secret-content-0",
		"yaq-keyvault-map-test-1": "secret-content-1",
	}
	var client *azsecrets.Client

	BeforeEach(func() {
		kvName = testutils.KeyvaultName()
		client = testutils.KVClient(kvName)
	})

	BeforeEach(func() {
		err := output.PushMap("keyvault-secrets:"+kvName, secretMap, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		for s := range secretMap {
			testutils.DeleteSecret(client, s)
		}
	})

	It("writes all the secrets", func() {
		for s, v := range secretMap {
			response, err := client.GetSecret(context.Background(), s, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(*response.Value).To(Equal(v))
		}
	})
})
