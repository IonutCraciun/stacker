for file in roots/*; do
btrfs property set -ts ${file} ro false;
done
