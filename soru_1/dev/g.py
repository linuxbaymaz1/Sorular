import paramiko
from datetime import datetime
import os
from dotenv import load_dotenv

load_dotenv(dotenv_path=os.path.join(os.path.dirname(__file__), ".env"))

HOST = os.getenv("SSH_HOST")
USER = os.getenv("SSH_USER")
PASS = os.getenv("SSH_PASS")
PORT = int(os.getenv("SSH_PORT", 22))

COMMANDS = [
    "uname -a",
    "whoami",
    "uptime",
    "ps aux",
    "df -h"
]


def connect():
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    ssh.connect(
        HOST,
        port=PORT,
        username=USER,
        password=PASS,
        timeout=5
    )

    return ssh


def run_commands(ssh):
    results = {}

    for cmd in COMMANDS:
        stdin, stdout, stderr = ssh.exec_command(cmd)

        out = stdout.read().decode("utf-8", errors="ignore")
        err = stderr.read().decode("utf-8", errors="ignore")

        results[cmd] = out + err

    return results


def search_keywords(results, keywords):
    matches = []

    for cmd, output in results.items():
        for line in output.splitlines():
            for kw in keywords:
                if kw.lower() in line.lower():
                    matches.append({
                        "command": cmd,
                        "keyword": kw,
                        "line": line.strip()
                    })

    return matches


def generate_report(results, matches, keywords):
    now = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    report = []
    report.append("# SSH ANALYSIS REPORT")
    report.append(f"Date: {now}\n")

    report.append("## Executed Commands")
    for c in COMMANDS:
        report.append(f"- {c}")
    report.append("")

    report.append("## Keywords")
    report.append(", ".join(keywords) + "\n")

    report.append("## Matches Found")

    if not matches:
        report.append("No matches found.\n")
    else:
        for i, m in enumerate(matches, 1):
            report.append(f"### Match {i}")
            report.append(f"- Command: {m['command']}")
            report.append(f"- Keyword: {m['keyword']}")
            report.append(f"- Line: {m['line']}\n")

    return "\n".join(report)


def save_report(text):
    with open("report.txt", "w", encoding="utf-8") as f:
        f.write(text)


def main():
    print("SSH analiz başlıyor...\n")

    raw = input("Aranacak keywordleri gir (virgül ile): ")
    keywords = [x.strip() for x in raw.split(",") if x.strip()]

    try:
        ssh = connect()
    except Exception as e:
        print("SSH bağlantı hatası:", e)
        return

    results = run_commands(ssh)
    matches = search_keywords(results, keywords)

    report = generate_report(results, matches, keywords)

    print("\n--- RAPOR ---\n")
    print(report)

    save_report(report)

    ssh.close()
    print("\nRapor report.txt olarak kaydedildi.")


if __name__ == "__main__":
    main()
