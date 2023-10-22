{
  description = "Web API with go";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    sqlboiler.url = "github:DGollings/nix-sqlboiler";
  };

  outputs = { self, nixpkgs, sqlboiler }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    with pkgs;
    {
      devShell.${system} = mkShell {
        buildInputs = [
          go_1_21
          postgresql_16
          oapi-codegen
          sqlboiler.packages.${system}.sqlboiler
          go-migrate
          air
          git-crypt
        ];

        shellHook = ''
          set -e
          export PORT=8080
          export PGUSER=postgres
          export PGDIR=$(pwd)/.postgres
          export PGHOST=$PGDIR
          export PGDATA=$PGDIR/data
          export PGLOG=$PGDIR/log
          export DATABASE_URL="postgresql:///postgres?host=$PGDIR&user=$PGUSER"

          if test ! -d "$PGDIR"; then
              mkdir "$PGDIR"
          fi

          if [ ! -d "$PGDATA" ]; then
              echo 'Initializing postgresql database...'
              initdb "$PGDATA" --auth=trust >/dev/null -U "$PGUSER"
          fi

          nix build .#postgres
        '';
      };

      packages.${system}.postgres = writeShellScriptBin "postgres.sh" ''
        set -e

        postgres_is_running() {
            pg_ctl status 2>/dev/null | grep "server is running" 1>/dev/null
        }

        postgres_stop() {
            echo 'stopping postgresql'
            pg_ctl stop
        }

        postgres_start() {
            if ! postgres_is_running; then
                echo 'starting postgresql'
                pg_ctl start \
                    -l "$PGLOG" \
                    -o "-c listen_addresses='0.0.0.0' -c unix_socket_directories=$PGDIR"
            else
                echo 'postgresql is already running'
            fi
            echo "connect to postgres with 'psql -U $PGUSER -d postgres'"
        }

        case "$1" in
        start)
          postgres_start
          ;;
        stop)
          postgres_stop
          ;;
        *)
          echo >&2 "Usage: $0 <start|stop>"
          exit 1
          ;;
        esac
      '';
    };
}
