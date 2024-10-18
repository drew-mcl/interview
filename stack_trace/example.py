import sys

def main():
    if len(sys.argv) != 2:
        print("Usage: python log_counter.py <logfile>")
        return

    log_file = sys.argv[1]
    try:
        total_errors, messages = count_error_logs(log_file)
    except Exception as e:
        print(f"Error processing log file: {e}")
        return

    print("===== Error Log Counts =====")
    print(f"Total ERROR logs: {total_errors}\n")
    print("===== Error Messages =====")
    for msg in messages:
        print(msg)

def is_error_line(line):
    return "ERROR" in line or "ERR" in line

def is_indented_line(line):
    return line.startswith((' ', '\t'))

def count_error_logs(filename):
    with open(filename, 'r') as file:
        lines = file.readlines()

    total_errors = 0
    messages = []
    i = 0

    while i < len(lines):
        line = lines[i].rstrip('\n')

        if is_error_line(line):
            total_errors += 1

            if "ERROR" in line:
                level = "ERROR"
            elif (" ERR " in line or line.startswith("ERR ") or line.endswith(" ERR")):
                level = "ERR"
            else:
                i += 1
                continue

            idx = line.rfind('-')
            if idx != -1:
                message = line[idx+1:].strip()
            else:
                message = line.strip()

            stack_trace_lines = []
            j = i + 1
            while j < len(lines) and is_indented_line(lines[j]):
                stack_trace_lines.append(lines[j].strip())
                j += 1

            formatted_message = f"{level}: {message}"
            if stack_trace_lines:
                formatted_message += "\n  Stack Trace:"
                for trace_line in stack_trace_lines:
                    formatted_message += f"\n    {trace_line}"

            messages.append(formatted_message)
            i = j
        else:
            i += 1

    return total_errors, messages

if __name__ == "__main__":
    main()