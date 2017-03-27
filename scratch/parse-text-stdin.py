import sys
import re
import time


def read_input(file):
    for line in file:
        yield line

def check_matches(line, regexes):
    match = False
    for name,r in regexes.items():
        m = re.search(r,line)
        if m:
            match = True
            '''
            if name == 'ipaddr':
                print("matched ipaddr: %s" % m.group())
            elif name == 'hostname':
                print("matched hostanme: %s" % m.group())
            '''
    return(match)


def main():
    start = time.time()
    re1 = re.compile('\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b')
    re2 = re.compile('hostname: [a-z0-9\-]*')
    matchcount = 0
    linecount = 0
    regexes = {'ipaddr':re1, 'hostname': re2}
    data = read_input(sys.stdin)
    for thing in data:
        linecount += 1
        matched = check_matches(thing, regexes)
        if matched:
            matchcount += 1
    print("lines: %d" % linecount)
    print("matches: %d" % matchcount)
    end = time.time()
    print("timetaken: %f ms" % ((end-start)*1000))


if __name__ == "__main__":
    main()
