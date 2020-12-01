with import <nixpkgs> { };
mkShell rec {
  buildInputs = [
    stdenv
    glibc.static
    cmake
  ];
  CFLAGS="-I${pkgs.glibc.dev}/include";
  LDFLAGS="-L${pkgs.glibc}/lib";
}
