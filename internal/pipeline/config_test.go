package pipeline_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/internal/pipeline"
)

var _ = Describe("Config", func() {
	DescribeTable("parse arguments successfully",
		func(args []string, i []string, d string, o string) {
			config, output, err := GetConfig("yaq", args)
			Expect(err).To(BeNil())
			Expect(output).To(Equal(""))
			Expect([]string(config.Input)).To(Equal(i))
			Expect(config.DumpTo).To(Equal(d))
			Expect(config.Output).To(Equal(o))
		},
		Entry("default arguments", []string{"-i", "stdin", "-d", "json", "-o", "stdout"}, []string{"stdin"}, "json", "stdout"),
		Entry("all arguments", []string{"-i", "the-input", "-d", "the-dump-format", "-o", "the-output"}, []string{"the-input"}, "the-dump-format", "the-output"),
		Entry("missing input argument", []string{"-d", "the-dump-format", "-o", "the-output"}, nil, "the-dump-format", "the-output"),
		Entry("missing dump argument", []string{"-o", "the-output", "-i", "the-input"}, []string{"the-input"}, "json", "the-output"),
		Entry("missing output argument", []string{"-i", "the-input", "-d", "the-dump-format"}, []string{"the-input"}, "the-dump-format", "stdout"),
		Entry("multipart argument", []string{"-i", "the-input:the-parameter"}, []string{"the-input:the-parameter"}, "json", "stdout"),
		Entry("repeated arguments", []string{"-i", "the-first-input", "-i", "the-second-input"}, []string{"the-first-input", "the-second-input"}, "json", "stdout"),
	)
	Context("unknown argument", func() {
		It("should error", func() {
			config, output, err := GetConfig("yaq", []string{"-i", "stdin", "-d", "json", "-x", "stdout"})
			Expect(err).To(HaveOccurred())
			Expect(output).To(ContainSubstring("flag provided but not defined: -x"))
			Expect(config).To(BeNil())
		})
	})
	Context("missing argument", func() {
		It("should error", func() {
			config, output, err := GetConfig("yaq", []string{"-i", "stdin", "-d", "json", "-o"})
			Expect(err).To(HaveOccurred())
			Expect(output).To(ContainSubstring("flag needs an argument: -o"))
			Expect(config).To(BeNil())
		})
	})
})
