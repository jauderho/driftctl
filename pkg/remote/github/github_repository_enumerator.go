package github

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	tf "github.com/cloudskiff/driftctl/pkg/remote/terraform"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/github"
)

type GithubRepositoryEnumerator struct {
	repository     GithubRepository
	factory        resource.ResourceFactory
	providerConfig tf.TerraformProviderConfig
}

func NewGithubRepositoryEnumerator(repo GithubRepository, factory resource.ResourceFactory, providerConfig tf.TerraformProviderConfig) *GithubRepositoryEnumerator {
	return &GithubRepositoryEnumerator{
		repository:     repo,
		factory:        factory,
		providerConfig: providerConfig,
	}
}

func (g *GithubRepositoryEnumerator) SupportedType() resource.ResourceType {
	return github.GithubRepositoryResourceType
}

func (g *GithubRepositoryEnumerator) Enumerate() ([]resource.Resource, error) {
	ids, err := g.repository.ListRepositories()
	if err != nil {
		return nil, remoteerror.NewResourceEnumerationError(err, string(g.SupportedType()))
	}

	results := make([]resource.Resource, len(ids))

	for _, id := range ids {
		results = append(
			results,
			g.factory.CreateAbstractResource(
				string(g.SupportedType()),
				id,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}