import os
import re

base_dir = "/home/kali/calisma/son/metrics"
os.makedirs(base_dir, exist_ok=True)

with open("/home/kali/calisma/massive_results.txt", "r") as f:
    lines = f.readlines()

current_project = None
project_content = []

for line in lines:
    # Match project headers like:
    # GO PROJESİ: gin
    # PYTHON PROJESİ: flask
    # VUE/NODE PROJESİ: vue
    match = re.match(r"(GO|PYTHON|VUE/NODE) PROJESİ:\s+(\S+)", line.strip())
    if match:
        # Save previous project
        if current_project and project_content:
            proj_dir = os.path.join(base_dir, current_project)
            os.makedirs(proj_dir, exist_ok=True)
            with open(os.path.join(proj_dir, "summary.txt"), "w") as out_f:
                out_f.writelines(project_content)
        
        current_project = match.group(2)
        project_content = [line]
    elif current_project:
        project_content.append(line)

# Save the last project
if current_project and project_content:
    proj_dir = os.path.join(base_dir, current_project)
    os.makedirs(proj_dir, exist_ok=True)
    with open(os.path.join(proj_dir, "summary.txt"), "w") as out_f:
        out_f.writelines(project_content)

print("Split completed.")
