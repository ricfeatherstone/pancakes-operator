package controllers

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pancakesv1alpha1 "github.com/ricfeatherstone/pancakes-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var _ = Describe("Batch Controller", func() {
	const (
		BatchName      = "example-batch"
		BatchNamespace = "default"

		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating a Batch", func() {
		It("Should be able to create", func() {
			By("Creating a new Batch")
			ctx := context.Background()
			batch := &pancakesv1alpha1.Batch{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      BatchName,
					Namespace: BatchNamespace,
				},
				Spec: pancakesv1alpha1.BatchSpec{
					Flour:        pancakesv1alpha1.Flour{Type: "Plain", Amount: 100, Weight: "Gram"},
					Eggs:         pancakesv1alpha1.Eggs{Number: 2, Size: "Large"},
					Milk:         pancakesv1alpha1.Milk{Amount: 300, Unit: "MilliLitre"},
					Oil:          pancakesv1alpha1.Oil{Type: "Vegetable", Amount: 1, Unit: "TableSpoon"},
					LemonToServe: pancakesv1alpha1.LemonWedges{Number: 4},
				},
			}

			Expect(k8sClient.Create(ctx, batch)).Should(Succeed())
		})
	})
})
