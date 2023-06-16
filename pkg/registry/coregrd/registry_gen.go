// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     GRDKindRegistryJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package coregrd

import (
	"context"
	"fmt"

	"github.com/grafana/dskit/services"
	kindsv1 "github.com/grafana/grafana-apiserver/pkg/apis/kinds/v1"
	applyConfig "github.com/grafana/grafana-apiserver/pkg/client/applyconfiguration/kinds/v1"
	grdClientset "github.com/grafana/grafana-apiserver/pkg/client/clientset/clientset/typed/kinds/v1"
	"github.com/grafana/grafana/pkg/modules"
	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/grafana/pkg/services/k8s/apiserver"
	"github.com/grafana/kindsys"
	"github.com/grafana/thema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// New constructs a new [Registry].
//
// All calling code within grafana/grafana is expected to use Grafana's
// singleton [thema.Runtime], returned from [cuectx.GrafanaThemaRuntime]. If nil
// is passed, the singleton will be used.
func New(
	rt *thema.Runtime,
	restConfigProvider apiserver.RestConfigProvider,
) *Registry {
	breg := corekind.NewBase(rt)
	r := doNewRegistry(
		breg,
		restConfigProvider,
	)
	r.BasicService = services.NewBasicService(r.start, r.run, nil).WithName(modules.KubernetesRegistration)
	return r
}

type Registry struct {
	*services.BasicService
	breg               *corekind.Base
	restConfigProvider apiserver.RestConfigProvider
}

func (r *Registry) start(ctx context.Context) error {
	clientSet, err := grdClientset.NewForConfig(r.restConfigProvider.GetRestConfig())
	if err != nil {
		return err
	}

	/************************ AccessPolicy ************************/
	accesspolicyGRD, err := r.getGRD(r.breg.AccessPolicy())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, accesspolicyGRD, metav1.CreateOptions{})

	accesspolicyApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range accesspolicyGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		accesspolicyApplyVersions = append(accesspolicyApplyVersions, version)
	}

	accesspolicyApplyNames := applyConfig.GrafanaKindNames().
		WithKind(accesspolicyGRD.Spec.Names.Kind).
		WithListKind(accesspolicyGRD.Spec.Names.ListKind).
		WithSingular(accesspolicyGRD.Spec.Names.Singular).
		WithPlural(accesspolicyGRD.Spec.Names.Plural).
		WithCategories(accesspolicyGRD.Spec.Names.Categories...).
		WithShortNames(accesspolicyGRD.Spec.Names.ShortNames...)

	accesspolicyApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(accesspolicyGRD.Spec.Group).
		WithNames(accesspolicyApplyNames).
		WithScope(accesspolicyGRD.Spec.Scope).
		WithVersions(accesspolicyApplyVersions...).
		WithPreserveUnknownFields(accesspolicyGRD.Spec.PreserveUnknownFields)

	accesspolicyApplyConfig := applyConfig.GrafanaKind(accesspolicyGRD.ObjectMeta.Name).
		WithSpec(accesspolicyApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, accesspolicyApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Dashboard ************************/
	dashboardGRD, err := r.getGRD(r.breg.Dashboard())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, dashboardGRD, metav1.CreateOptions{})

	dashboardApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range dashboardGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		dashboardApplyVersions = append(dashboardApplyVersions, version)
	}

	dashboardApplyNames := applyConfig.GrafanaKindNames().
		WithKind(dashboardGRD.Spec.Names.Kind).
		WithListKind(dashboardGRD.Spec.Names.ListKind).
		WithSingular(dashboardGRD.Spec.Names.Singular).
		WithPlural(dashboardGRD.Spec.Names.Plural).
		WithCategories(dashboardGRD.Spec.Names.Categories...).
		WithShortNames(dashboardGRD.Spec.Names.ShortNames...)

	dashboardApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(dashboardGRD.Spec.Group).
		WithNames(dashboardApplyNames).
		WithScope(dashboardGRD.Spec.Scope).
		WithVersions(dashboardApplyVersions...).
		WithPreserveUnknownFields(dashboardGRD.Spec.PreserveUnknownFields)

	dashboardApplyConfig := applyConfig.GrafanaKind(dashboardGRD.ObjectMeta.Name).
		WithSpec(dashboardApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, dashboardApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Folder ************************/
	folderGRD, err := r.getGRD(r.breg.Folder())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, folderGRD, metav1.CreateOptions{})

	folderApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range folderGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		folderApplyVersions = append(folderApplyVersions, version)
	}

	folderApplyNames := applyConfig.GrafanaKindNames().
		WithKind(folderGRD.Spec.Names.Kind).
		WithListKind(folderGRD.Spec.Names.ListKind).
		WithSingular(folderGRD.Spec.Names.Singular).
		WithPlural(folderGRD.Spec.Names.Plural).
		WithCategories(folderGRD.Spec.Names.Categories...).
		WithShortNames(folderGRD.Spec.Names.ShortNames...)

	folderApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(folderGRD.Spec.Group).
		WithNames(folderApplyNames).
		WithScope(folderGRD.Spec.Scope).
		WithVersions(folderApplyVersions...).
		WithPreserveUnknownFields(folderGRD.Spec.PreserveUnknownFields)

	folderApplyConfig := applyConfig.GrafanaKind(folderGRD.ObjectMeta.Name).
		WithSpec(folderApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, folderApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ LibraryPanel ************************/
	librarypanelGRD, err := r.getGRD(r.breg.LibraryPanel())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, librarypanelGRD, metav1.CreateOptions{})

	librarypanelApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range librarypanelGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		librarypanelApplyVersions = append(librarypanelApplyVersions, version)
	}

	librarypanelApplyNames := applyConfig.GrafanaKindNames().
		WithKind(librarypanelGRD.Spec.Names.Kind).
		WithListKind(librarypanelGRD.Spec.Names.ListKind).
		WithSingular(librarypanelGRD.Spec.Names.Singular).
		WithPlural(librarypanelGRD.Spec.Names.Plural).
		WithCategories(librarypanelGRD.Spec.Names.Categories...).
		WithShortNames(librarypanelGRD.Spec.Names.ShortNames...)

	librarypanelApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(librarypanelGRD.Spec.Group).
		WithNames(librarypanelApplyNames).
		WithScope(librarypanelGRD.Spec.Scope).
		WithVersions(librarypanelApplyVersions...).
		WithPreserveUnknownFields(librarypanelGRD.Spec.PreserveUnknownFields)

	librarypanelApplyConfig := applyConfig.GrafanaKind(librarypanelGRD.ObjectMeta.Name).
		WithSpec(librarypanelApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, librarypanelApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Playlist ************************/
	playlistGRD, err := r.getGRD(r.breg.Playlist())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, playlistGRD, metav1.CreateOptions{})

	playlistApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range playlistGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		playlistApplyVersions = append(playlistApplyVersions, version)
	}

	playlistApplyNames := applyConfig.GrafanaKindNames().
		WithKind(playlistGRD.Spec.Names.Kind).
		WithListKind(playlistGRD.Spec.Names.ListKind).
		WithSingular(playlistGRD.Spec.Names.Singular).
		WithPlural(playlistGRD.Spec.Names.Plural).
		WithCategories(playlistGRD.Spec.Names.Categories...).
		WithShortNames(playlistGRD.Spec.Names.ShortNames...)

	playlistApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(playlistGRD.Spec.Group).
		WithNames(playlistApplyNames).
		WithScope(playlistGRD.Spec.Scope).
		WithVersions(playlistApplyVersions...).
		WithPreserveUnknownFields(playlistGRD.Spec.PreserveUnknownFields)

	playlistApplyConfig := applyConfig.GrafanaKind(playlistGRD.ObjectMeta.Name).
		WithSpec(playlistApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, playlistApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Preferences ************************/
	preferencesGRD, err := r.getGRD(r.breg.Preferences())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, preferencesGRD, metav1.CreateOptions{})

	preferencesApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range preferencesGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		preferencesApplyVersions = append(preferencesApplyVersions, version)
	}

	preferencesApplyNames := applyConfig.GrafanaKindNames().
		WithKind(preferencesGRD.Spec.Names.Kind).
		WithListKind(preferencesGRD.Spec.Names.ListKind).
		WithSingular(preferencesGRD.Spec.Names.Singular).
		WithPlural(preferencesGRD.Spec.Names.Plural).
		WithCategories(preferencesGRD.Spec.Names.Categories...).
		WithShortNames(preferencesGRD.Spec.Names.ShortNames...)

	preferencesApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(preferencesGRD.Spec.Group).
		WithNames(preferencesApplyNames).
		WithScope(preferencesGRD.Spec.Scope).
		WithVersions(preferencesApplyVersions...).
		WithPreserveUnknownFields(preferencesGRD.Spec.PreserveUnknownFields)

	preferencesApplyConfig := applyConfig.GrafanaKind(preferencesGRD.ObjectMeta.Name).
		WithSpec(preferencesApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, preferencesApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ PublicDashboard ************************/
	publicdashboardGRD, err := r.getGRD(r.breg.PublicDashboard())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, publicdashboardGRD, metav1.CreateOptions{})

	publicdashboardApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range publicdashboardGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		publicdashboardApplyVersions = append(publicdashboardApplyVersions, version)
	}

	publicdashboardApplyNames := applyConfig.GrafanaKindNames().
		WithKind(publicdashboardGRD.Spec.Names.Kind).
		WithListKind(publicdashboardGRD.Spec.Names.ListKind).
		WithSingular(publicdashboardGRD.Spec.Names.Singular).
		WithPlural(publicdashboardGRD.Spec.Names.Plural).
		WithCategories(publicdashboardGRD.Spec.Names.Categories...).
		WithShortNames(publicdashboardGRD.Spec.Names.ShortNames...)

	publicdashboardApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(publicdashboardGRD.Spec.Group).
		WithNames(publicdashboardApplyNames).
		WithScope(publicdashboardGRD.Spec.Scope).
		WithVersions(publicdashboardApplyVersions...).
		WithPreserveUnknownFields(publicdashboardGRD.Spec.PreserveUnknownFields)

	publicdashboardApplyConfig := applyConfig.GrafanaKind(publicdashboardGRD.ObjectMeta.Name).
		WithSpec(publicdashboardApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, publicdashboardApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Role ************************/
	roleGRD, err := r.getGRD(r.breg.Role())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, roleGRD, metav1.CreateOptions{})

	roleApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range roleGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		roleApplyVersions = append(roleApplyVersions, version)
	}

	roleApplyNames := applyConfig.GrafanaKindNames().
		WithKind(roleGRD.Spec.Names.Kind).
		WithListKind(roleGRD.Spec.Names.ListKind).
		WithSingular(roleGRD.Spec.Names.Singular).
		WithPlural(roleGRD.Spec.Names.Plural).
		WithCategories(roleGRD.Spec.Names.Categories...).
		WithShortNames(roleGRD.Spec.Names.ShortNames...)

	roleApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(roleGRD.Spec.Group).
		WithNames(roleApplyNames).
		WithScope(roleGRD.Spec.Scope).
		WithVersions(roleApplyVersions...).
		WithPreserveUnknownFields(roleGRD.Spec.PreserveUnknownFields)

	roleApplyConfig := applyConfig.GrafanaKind(roleGRD.ObjectMeta.Name).
		WithSpec(roleApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, roleApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ RoleBinding ************************/
	rolebindingGRD, err := r.getGRD(r.breg.RoleBinding())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, rolebindingGRD, metav1.CreateOptions{})

	rolebindingApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range rolebindingGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		rolebindingApplyVersions = append(rolebindingApplyVersions, version)
	}

	rolebindingApplyNames := applyConfig.GrafanaKindNames().
		WithKind(rolebindingGRD.Spec.Names.Kind).
		WithListKind(rolebindingGRD.Spec.Names.ListKind).
		WithSingular(rolebindingGRD.Spec.Names.Singular).
		WithPlural(rolebindingGRD.Spec.Names.Plural).
		WithCategories(rolebindingGRD.Spec.Names.Categories...).
		WithShortNames(rolebindingGRD.Spec.Names.ShortNames...)

	rolebindingApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(rolebindingGRD.Spec.Group).
		WithNames(rolebindingApplyNames).
		WithScope(rolebindingGRD.Spec.Scope).
		WithVersions(rolebindingApplyVersions...).
		WithPreserveUnknownFields(rolebindingGRD.Spec.PreserveUnknownFields)

	rolebindingApplyConfig := applyConfig.GrafanaKind(rolebindingGRD.ObjectMeta.Name).
		WithSpec(rolebindingApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, rolebindingApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	/************************ Team ************************/
	teamGRD, err := r.getGRD(r.breg.Team())
	if err != nil {
		return err
	}

	//_, _ = clientSet.GrafanaKinds().Create(ctx, teamGRD, metav1.CreateOptions{})

	teamApplyVersions := make([]*applyConfig.GrafanaKindVersionApplyConfiguration, 0)
	for _, v := range teamGRD.Spec.Versions {
		subresource := applyConfig.GrafanaResourceSubresources()
		if v.Subresources != nil && v.Subresources.Status != nil {
			subresource = subresource.WithStatus(*v.Subresources.Status)
		}

		if v.Subresources != nil && v.Subresources.Ref != nil {
			subresource = subresource.WithRef(*v.Subresources.Ref)
		}

		if v.Subresources != nil && v.Subresources.History != nil {
			subresource = subresource.WithHistory(*v.Subresources.History)
		}

		version := applyConfig.GrafanaKindVersion().
			WithName(v.Name).
			WithServed(v.Served).
			WithStorage(v.Storage).
			WithDeprecated(v.Deprecated).
			WithSubresources(subresource)
		if v.DeprecationWarning != nil {
			version = version.WithDeprecationWarning(*v.DeprecationWarning)
		}
		teamApplyVersions = append(teamApplyVersions, version)
	}

	teamApplyNames := applyConfig.GrafanaKindNames().
		WithKind(teamGRD.Spec.Names.Kind).
		WithListKind(teamGRD.Spec.Names.ListKind).
		WithSingular(teamGRD.Spec.Names.Singular).
		WithPlural(teamGRD.Spec.Names.Plural).
		WithCategories(teamGRD.Spec.Names.Categories...).
		WithShortNames(teamGRD.Spec.Names.ShortNames...)

	teamApplySpec := applyConfig.GrafanaKindSpec().
		WithGroup(teamGRD.Spec.Group).
		WithNames(teamApplyNames).
		WithScope(teamGRD.Spec.Scope).
		WithVersions(teamApplyVersions...).
		WithPreserveUnknownFields(teamGRD.Spec.PreserveUnknownFields)

	teamApplyConfig := applyConfig.GrafanaKind(teamGRD.ObjectMeta.Name).
		WithSpec(teamApplySpec)

	_, err = clientSet.GrafanaKinds().Apply(ctx, teamApplyConfig, metav1.ApplyOptions{FieldManager: "grafana"})
	if err != nil {
		return err
	}

	return nil
}

func (r *Registry) run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (r *Registry) getGRD(k kindsys.Kind) (*kindsv1.GrafanaKind, error) {
	kind, is := k.(kindsys.Core)
	if !is {
		return nil, nil
	}

	props := kind.Def().Properties
	lin := kind.Lineage()

	// We need to go through every schema, as they all have to be defined in the CRD
	sch, err := lin.Schema(thema.SV(0, 0))
	if err != nil {
		return nil, err
	}

	resource := kindsv1.GrafanaKind{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kinds.grafana.com/v1",
			Kind:       "GrafanaKind",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.%s", props.PluralMachineName, "core.kinds.grafana.com"),
		},
		Spec: kindsv1.GrafanaKindSpec{
			Group: "core.kinds.grafana.com",
			Scope: "Namespaced",
			Names: kindsv1.GrafanaKindNames{
				Kind:     props.Name,
				ListKind: props.Name + "List",
				Singular: props.MachineName,
				Plural:   props.PluralMachineName,
			},
			Versions: make([]kindsv1.GrafanaKindVersion, 0),
		},
	}
	latest := lin.Latest().Version()

	for sch != nil {
		vstr := versionString(sch.Version())
		if props.Maturity.Less(kindsys.MaturityStable) {
			vstr = "v0-alpha"
		}

		ver := kindsv1.GrafanaKindVersion{
			Name:       vstr,
			Served:     true,
			Storage:    sch.Version() == latest,
			Deprecated: false,
			Subresources: &kindsv1.GrafanaResourceSubresources{
				Status:  &kindsv1.GrafanaResourceSubresourceStatus{},
				History: &kindsv1.GrafanaResourceSubresourceHistory{},
				Ref:     &kindsv1.GrafanaResourceSubresourceRef{},
			},
		}

		resource.Spec.Versions = append(resource.Spec.Versions, ver)
		sch = sch.Successor()
	}

	return &resource, nil
}

func versionString(version thema.SyntacticVersion) string {
	// TODO: v0.0-alpha throws the apiservice registration off in aggregated mode
	// Cannot use dot in the DNS subdomain prefix in front of *.kinds.grafana.com
	return fmt.Sprintf("v%d", version[0]) // , version[1]
}

func doNewRegistry(
	breg *corekind.Base,
	restConfigProvider apiserver.RestConfigProvider,
) *Registry {
	return &Registry{
		breg:               breg,
		restConfigProvider: restConfigProvider,
	}
}
