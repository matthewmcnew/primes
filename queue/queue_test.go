package queue_test

import (
	. "github.com/matthewmcnew/primes/queue"
	"container/heap"
	"github.com/matthewmcnew/primes/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	Describe("Pop", func() {
		It("Removes the next element", func() {
			queue := &Queue{}
			heap.Init(queue)

			item := &models.CalculatedResult{Prime: 1, Job: 1}
			heap.Push(queue, item)

			Otheritem := &models.CalculatedResult{Prime: 1, Job: 5}
			heap.Push(queue, Otheritem)

			retrievedItem := heap.Pop(queue).(*models.CalculatedResult)

			Expect(retrievedItem).To(Equal(item))
		})
	})

	Describe("LowestJob", func() {
		It("Returns the next wihtout returning", func() {
			queue := &Queue{}
			heap.Init(queue)

			item := &models.CalculatedResult{Prime: 1, Job: 1}
			heap.Push(queue, item)

			Otheritem := &models.CalculatedResult{Prime: 1, Job: 5}
			heap.Push(queue, Otheritem)

			lowestValue := queue.LowestJob()

			Expect(lowestValue).To(Equal(1))
		})
	})

	Describe("Len", func() {
			It("Returns the Length of the queue", func() {
					queue := &Queue{}
					heap.Init(queue)

					Expect(queue.Len()).To(Equal(0))

					item := &models.CalculatedResult{Prime: 1, Job: 1}
					heap.Push(queue, item)

					Expect(queue.Len()).To(Equal(1))

					Otheritem := &models.CalculatedResult{Prime: 1, Job: 5}
					heap.Push(queue, Otheritem)

					Expect(queue.Len()).To(Equal(2))
				})
		})
})
