package v7action_test

import (
	"errors"

	. "code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/actor/v7action/v7actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Space Quota Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v7actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		actor, fakeCloudControllerClient, _, _, _, _ = NewTestActor()

	})

	Describe("CreateSpaceQuota", func() {
		var (
			spaceQuotaName   string
			organizationGuid string
			warnings         Warnings
			executeErr       error
			limits           QuotaLimits
		)

		JustBeforeEach(func() {
			warnings, executeErr = actor.CreateSpaceQuota(spaceQuotaName, organizationGuid, limits)
		})

		When("creating a space quota with all values set", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.CreateSpaceQuotaReturns(ccv3.SpaceQuota{}, ccv3.Warnings{"some-quota-warning"}, nil)
				limits = QuotaLimits{
					TotalMemoryInMB:       types.NullInt{IsSet: true, Value: 2},
					PerProcessMemoryInMB:  types.NullInt{IsSet: true, Value: 3},
					TotalInstances:        types.NullInt{IsSet: true, Value: 4},
					PaidServicesAllowed:   true,
					TotalServiceInstances: types.NullInt{IsSet: true, Value: 6},
					TotalRoutes:           types.NullInt{IsSet: true, Value: 8},
					TotalReservedPorts:    types.NullInt{IsSet: true, Value: 9},
				}
			})

			It("makes the space quota", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(fakeCloudControllerClient.CreateSpaceQuotaCallCount()).To(Equal(1))
				givenSpaceQuota := fakeCloudControllerClient.CreateSpaceQuotaArgsForCall(0)

				Expect(givenSpaceQuota).To(Equal(ccv3.SpaceQuota{
					Name:    spaceQuotaName,
					OrgGUID: organizationGuid,
					Apps: ccv3.AppLimit{
						TotalMemory:       types.NullInt{IsSet: true, Value: 2},
						InstanceMemory:    types.NullInt{IsSet: true, Value: 3},
						TotalAppInstances: types.NullInt{IsSet: true, Value: 4},
					},
					Services: ccv3.ServiceLimit{
						TotalServiceInstances: types.NullInt{IsSet: true, Value: 6},
						PaidServicePlans:      true,
					},
					Routes: ccv3.RouteLimit{
						TotalRoutes:        types.NullInt{IsSet: true, Value: 8},
						TotalReservedPorts: types.NullInt{IsSet: true, Value: 9},
					},
					SpaceGUIDs: nil,
				}))

				Expect(warnings).To(ConsistOf("some-quota-warning"))
			})
		})

		When("creating a space quota with empty limits", func() {
			var (
				ccv3Quota ccv3.SpaceQuota
			)

			BeforeEach(func() {
				spaceQuotaName = "quota-name"
				limits = QuotaLimits{}

				ccv3Quota = ccv3.SpaceQuota{
					Name: spaceQuotaName,
					Apps: ccv3.AppLimit{
						TotalMemory:       types.NullInt{Value: 0, IsSet: true},
						InstanceMemory:    types.NullInt{Value: 0, IsSet: false},
						TotalAppInstances: types.NullInt{Value: 0, IsSet: false},
					},
					Services: ccv3.ServiceLimit{
						TotalServiceInstances: types.NullInt{Value: 0, IsSet: true},
						PaidServicePlans:      false,
					},
					Routes: ccv3.RouteLimit{
						TotalRoutes:        types.NullInt{Value: 0, IsSet: true},
						TotalReservedPorts: types.NullInt{Value: 0, IsSet: true},
					},
				}
				fakeCloudControllerClient.CreateSpaceQuotaReturns(
					ccv3Quota,
					ccv3.Warnings{"some-quota-warning"},
					nil,
				)
			})

			It("call the create endpoint with the respective values and returns warnings", func() {
				Expect(fakeCloudControllerClient.CreateSpaceQuotaCallCount()).To(Equal(1))

				Expect(warnings).To(ConsistOf("some-quota-warning"))

				passedQuota := fakeCloudControllerClient.CreateSpaceQuotaArgsForCall(0)
				Expect(passedQuota).To(Equal(ccv3Quota))
			})
		})

		When("creating a quota with all values set to unlimited", func() {
			var (
				ccv3Quota ccv3.SpaceQuota
			)

			BeforeEach(func() {
				spaceQuotaName = "quota-name"
				limits = QuotaLimits{
					TotalMemoryInMB:       types.NullInt{Value: -1, IsSet: true},
					PerProcessMemoryInMB:  types.NullInt{Value: -1, IsSet: true},
					TotalInstances:        types.NullInt{Value: -1, IsSet: true},
					TotalServiceInstances: types.NullInt{Value: -1, IsSet: true},
					TotalRoutes:           types.NullInt{Value: -1, IsSet: true},
					TotalReservedPorts:    types.NullInt{Value: -1, IsSet: true},
				}
				ccv3Quota = ccv3.SpaceQuota{
					Name: spaceQuotaName,
					Apps: ccv3.AppLimit{
						TotalMemory:       types.NullInt{Value: -1, IsSet: false},
						InstanceMemory:    types.NullInt{Value: -1, IsSet: false},
						TotalAppInstances: types.NullInt{Value: -1, IsSet: false},
					},
					Services: ccv3.ServiceLimit{
						TotalServiceInstances: types.NullInt{Value: -1, IsSet: false},
						PaidServicePlans:      false,
					},
					Routes: ccv3.RouteLimit{
						TotalRoutes:        types.NullInt{Value: -1, IsSet: false},
						TotalReservedPorts: types.NullInt{Value: -1, IsSet: false},
					},
				}
				fakeCloudControllerClient.CreateSpaceQuotaReturns(
					ccv3Quota,
					ccv3.Warnings{"some-quota-warning"},
					nil,
				)
			})

			It("call the create endpoint with the respective values and returns warnings", func() {
				Expect(fakeCloudControllerClient.CreateSpaceQuotaCallCount()).To(Equal(1))

				Expect(warnings).To(ConsistOf("some-quota-warning"))

				passedQuota := fakeCloudControllerClient.CreateSpaceQuotaArgsForCall(0)
				Expect(passedQuota).To(Equal(ccv3Quota))
			})
		})

		When("creating a quota returns an error", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.CreateSpaceQuotaReturns(
					ccv3.SpaceQuota{},
					ccv3.Warnings{"some-quota-warning"},
					errors.New("create-error"),
				)
			})

			It("returns the error and warnings", func() {
				Expect(fakeCloudControllerClient.CreateSpaceQuotaCallCount()).To(Equal(1))

				Expect(warnings).To(ConsistOf("some-quota-warning"))
				Expect(executeErr).To(MatchError("create-error"))
			})
		})
	})
})