/*
Copyright 2019-2023 TensorChord Inc.

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by informer-gen. DO NOT EDIT.

package modelzetes

import (
	internalinterfaces "github.com/tensorchord/openmodelz/modelzetes/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/tensorchord/openmodelz/modelzetes/pkg/client/informers/externalversions/modelzetes/v2alpha1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V2alpha1 provides access to shared informers for resources in V2alpha1.
	V2alpha1() v2alpha1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V2alpha1 returns a new v2alpha1.Interface.
func (g *group) V2alpha1() v2alpha1.Interface {
	return v2alpha1.New(g.factory, g.namespace, g.tweakListOptions)
}