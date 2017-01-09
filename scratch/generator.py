
import random
import sys


class Timestamp ():
    def __init__(self):
        self.year = random.randint(2007,2017)
        self.month = random.randint(1,12)
        self.day = random.randint(1,31)
        self.hour = random.randint(0,23)
        self.minute = random.randint(0,59)
        self.second = random.randint(0,59)
        self.millisecond = random.randint(0,1000)
    def dumpself(self):
        return self.__repr__()
    def __repr__(self):
        # sample format for RFC3339: 2017-01-09T16:15:09-05:00
        fmt = ("{0:04d}-{1:02d}-{2:02d}T{3:02d}:{4:02d}:{5:02d}-05:00")
        props = []
        props.append(self.year)
        props.append(self.month)
        props.append(self.day)
        props.append(self.hour)
        props.append(self.minute)
        props.append(self.second)
        # props.append(self.millisecond)
        return fmt.format(*props)
        #return()

class Ip ():
    def __init__(self):
        self.one = random.randint(1,255)
        self.two = random.randint(1,255)
        self.three = random.randint(1,255)
        self.four = random.randint(1,255)
    def __repr__(self):
        return("%s.%s.%s.%s" % (str(self.one),str(self.two),str(self.three),str(self.four)))
    def dumpself(self):
        return self.__repr__()
    

for i in range(int(sys.argv[1])):
    # every hundred lines throw a bogus non matching line
    if i % 100 == 0:
        print("999-89-asdf hosty not matchy no not cool")
        continue
    else:
        x = Ip()
        y = Timestamp()
        print("%s - machine with ip of %s and a hostname: haha-nono" % (y.dumpself(), x.dumpself()))
    







