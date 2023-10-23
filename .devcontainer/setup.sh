#!/bin/bash
set -ex
# Set up global gitignore for direnv.
mkdir -p ~/.config/git && printf '.direnv/\\n.envrc\\n' > ~/.config/git/ignore && git config --global core.excludesfile ~/.config/git/ignore
# Install git-crypt for codespace hook?
nix profile install nixpkgs#git-crypt
# Install, set up and allow direnv in workspace.
nix profile install nixpkgs#direnv nixpkgs#nix-direnv && mkdir -p ~/.config/direnv && echo 'source $HOME/.nix-profile/share/nix-direnv/direnvrc' >> ~/.config/direnv/direnvrc && cp envrc.recommended .envrc && direnv allow
# Add shell hook for direnv
echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
