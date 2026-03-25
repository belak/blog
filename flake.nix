{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{
      flake-parts,
      treefmt-nix,
      nixpkgs,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        treefmt-nix.flakeModule
      ];

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
          treefmt = {
            programs.nixfmt.enable = true;

            programs.goimports = {
              enable = true;
              excludes = [
                "*.sql.go"
                "*_templ.go"
              ];
            };

            programs.templ = {
              enable = true;
              package = config.packages.templ;
            };

            programs.prettier = {
              enable = true;
              includes = [
                "*.css"
                "*.md"
              ];
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
            default = pkgs.myBuildGoModule rec {
              pname = "belak-blog";
              version = "0.0.1";
              src = ./.;
              subPackages = [
                "cmd/belak-blog"
              ];
              vendorHash = "";
            };
          }
          // (lib.packagesFromDirectoryRecursive {
            callPackage = pkgs.callPackage;
            directory = ./nix/pkgs;
          });
        };
    };
}
