{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };

  };

  outputs =
    inputs@{
      flake-parts,
      nixpkgs,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {

      systems = [
        "aarch64-linux"
        "x86_64-linux"
        "aarch64-darwin"
        "x86_64-darwin"
      ];

      perSystem =
        {
          lib,
          pkgs,
          system,
          config,
          ...
        }:
        {
          formatter = pkgs.treefmt.withConfig {
            runtimeInputs = with pkgs; [
              nixfmt
              gotools
              config.packages.templ
              prettier
            ];

            settings = {
              formatter.nixfmt = {
                command = "nixfmt";
                includes = [ "*.nix" ];
              };

              formatter.goimports = {
                command = "goimports";
                includes = [ "*.go" ];
                excludes = [
                  "*.sql.go"
                  "*_templ.go"
                ];
              };

              formatter.templ = {
                command = "templ";
                options = [ "fmt" ];
                includes = [ "*.templ" ];
              };

              formatter.prettier = {
                command = "prettier";
                options = [ "--write" ];
                includes = [
                  "*.css"
                  "*.md"
                ];
                excludes = [
                  "**/chroma.css"
                  "**/simple-*.css"
                ];
              };
            };
          };

          devShells.default = pkgs.mkShell {
            packages = with pkgs; [
              air
              go
              gopls
              golangci-lint
              prettier

              config.packages.templ
            ];
          };

          packages = {
            default = pkgs.buildGoModule rec {
              pname = "belak-blog";
              version = "0.0.1";
              src = ./.;
              subPackages = [
                "cmd/belak-blog"
              ];
              vendorHash = "sha256-q95d5DmXFDELfkg3/oL0tF0qZDrXs0JjIZXTNA3fWtY=";
            };
          }
          // (lib.packagesFromDirectoryRecursive {
            callPackage = pkgs.callPackage;
            directory = ./nix/pkgs;
          });
        };
    };
}
