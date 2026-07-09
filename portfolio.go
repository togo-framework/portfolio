// Package portfolio is the Portfolio content pack for togo: Project resource + portfolio list & slug detail pages (with GitHub cards).
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package portfolio

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("portfolio", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
