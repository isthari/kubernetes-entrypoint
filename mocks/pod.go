package mocks

import (
	"fmt"

	v1core "k8s.io/client-go/1.5/kubernetes/typed/core/v1"
	api "k8s.io/client-go/1.5/pkg/api"
	v1 "k8s.io/client-go/1.5/pkg/api/v1"
	policy "k8s.io/client-go/1.5/pkg/apis/policy/v1alpha1"
	"k8s.io/client-go/1.5/pkg/watch"
	"k8s.io/client-go/1.5/rest"
)

type pClient struct {
}

func (p pClient) Get(name string) (*v1.Pod, error) {
	return &v1.Pod{
		ObjectMeta: v1.ObjectMeta{Name: name},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					Name:  "container_test",
					Ready: true,
				},
			},
			HostIP: "127.0.0.1",
		},
	}, nil

}
func (p pClient) Create(pod *v1.Pod) (*v1.Pod, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (p pClient) Delete(name string, options *api.DeleteOptions) error {
	return fmt.Errorf("Not implemented")
}

func (p pClient) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return fmt.Errorf("Not implemented")
}

func (p pClient) List(options api.ListOptions) (*v1.PodList, error) {
	return &v1.PodList{
		Items: []v1.Pod{
			{
				ObjectMeta: v1.ObjectMeta{Name: "podList"},
				Status: v1.PodStatus{
					HostIP: "127.0.01",
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: "True",
						},
					},
					ContainerStatuses: []v1.ContainerStatus{
						{
							Name:  "container_test",
							Ready: true,
						},
					},
				},
			},
		},
	}, nil

}

func (p pClient) Update(pod *v1.Pod) (*v1.Pod, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (p pClient) UpdateStatus(pod *v1.Pod) (*v1.Pod, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (p pClient) Watch(options api.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (p pClient) Bind(binding *v1.Binding) error {
	return fmt.Errorf("Not implemented")
}

func (p pClient) Evict(eviction *policy.Eviction) error {
	return fmt.Errorf("Not implemented")
}

func (p pClient) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	return nil
}

func (p pClient) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1.Pod, err error) {
	return nil, fmt.Errorf("Not implemented")
}
func NewPClient() v1core.PodInterface {
	return pClient{}
}
