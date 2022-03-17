import sys
def wc(filenames):
    print(filenames)
    results = {}
    for filename in filenames:
        chars = 0
        words = 0
        lines = 0
        try:
            with open(filename) as fh:
                for line in fh:
                    lines += 1
                    words += len(line.split())
                    chars += len(line)
            results[filename] = {
                'chars' : chars,
                'words' : words,
                'lines' : lines
            }
        except Exception as err:
            print(err)
        return results
        

if __name__ == '__main__':
    if len(sys.argv) < 1:
        exit("File not empty") 
    totals = {
        'lines': 0,
        'words': 0,
        'chars': 0,
    }

    results = wc(sys.argv[1:])
    for key in results:
        print(key,results)
        data = results[key]
        print("{} {} {} {}".format(data['lines'], data['words'], data['chars'], key))

        totals['lines'] = totals['lines'] + data['lines']
        totals['words'] = totals['words'] + data['words']
        totals['chars'] = totals['chars'] + data['chars']
    print("{} {} {} {}".format(totals['lines'], totals['words'], totals['chars'], 'total'))


