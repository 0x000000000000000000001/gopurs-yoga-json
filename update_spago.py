import yaml
import glob
import os

with open('spago.go.yaml', 'r') as f:
    data = yaml.safe_load(f)

for d in glob.glob('../gopurs-*'):
    if not os.path.isdir(d): continue
    pkg_name = os.path.basename(d).replace('gopurs-', '')
    if pkg_name == 'yoga-json': continue
    data['workspace']['extraPackages'][pkg_name] = {'path': d}

with open('spago.go.yaml', 'w') as f:
    yaml.dump(data, f, default_flow_style=False, sort_keys=True)
