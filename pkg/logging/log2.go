package logging

import (
    "runtime"
    "path/filepath"
    "fmt"
)

func Debug2(v ...interface{}) {
    setPrefix2(DEBUG)
    logger.Println(v)
}

func Info2(v ...interface{}) {
    setPrefix2(INFO)
    logger.Println(v)
}

func Warn2(v ...interface{}) {
    setPrefix2(WARNING)
    logger.Println(v)
}

func Error2(v ...interface{}) {
    setPrefix2(ERROR)
    logger.Println(v)
}

func Fatal2(v ...interface{}) {
    setPrefix2(FATAL)
    logger.Fatalln(v)
}

func setPrefix2(level Level) {
    _, file, line, ok := runtime.Caller(DefaultCallerDepth)
    if ok {
        logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
    } else {
        logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
    }

    logger.SetPrefix(logPrefix)
}
