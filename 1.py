import sys

def main():
    for line in sys.stdin:
        line = line.strip()
        print('"%s" => "%s",' % (line, line))
    return

if __name__ == '__main__':
    main()