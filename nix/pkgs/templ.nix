{
  lib,
  buildGoModule,
  fetchFromGitHub,
}:
buildGoModule rec {
  pname = "templ";
  version = "0.3.1001";

  src = fetchFromGitHub {
    owner = "a-h";
    repo = "templ";
    rev = "v${version}";
    hash = "sha256-146QxN+osvlzp8NTGm5TN2yvbu3cOodXfIVeIKsS+7I=";
  };

  vendorHash = "sha256-pVZjZCXT/xhBCMyZdR7kEmB9jqhTwRISFp63bQf6w5A=";

  env = {
    CGO_ENABLED = 0;
  };

  flags = [
    "-trimpath"
  ];

  ldflags = [
    "-s"
    "-w"
    "-extldflags -static"
  ];

  subPackages = [ "cmd/templ" ];
}
