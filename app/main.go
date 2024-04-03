package main

import (
  "fmt"
  "io"
  "log"
  "log/slog"
  "math/rand/v2"
  "os"
  "time"

  "github.com/urfave/cli/v2"
)

var (
  logFile = os.Getenv("HOME") + "/logs/app.log"
)

func main() {
  app := &cli.App{
    Name: "generate-logs",
    Usage: "This utility allows you to generate logs ",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name: "name",
        Aliases: []string{"n"},
        Usage: "Specify the app name. (required)",
        Required: true,
      },
      &cli.StringFlag{
        Name: "format",
        Aliases: []string{"f"},
        Value: "text",
        Usage: "Specify the format of the generated logs. The allowed values are \"text\" or \"json\".",
        Action: validateFormat,
      },
      &cli.StringFlag{
        Name: "output",
        Aliases: []string{"o"},
        Value: "stdout",
        Usage: fmt.Sprintf("Specify the output type. The allowed values are \"stdout\" or \"file\". If file is chosen, the logs will be generated in the following file: \"%s\".", logFile),
        Action: validateOutput, 
      },
      &cli.IntFlag{
        Name: "interval",
        Aliases: []string{"i"},
        Value: 5, 
        Usage: "Indicate the interval between two generated logs. Time expressed in seconds. (min: 1 - max: 10)",
        Action: validateInterval,
      },
    },
    Action: func (ctx *cli.Context) error {
      defer generateLogs(
        ctx.String("name"),
        ctx.String("format"),
        ctx.String("output"),
        ctx.Int("interval"),
      )

      return nil
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}

func validateFormat(ctx *cli.Context, v string) error {
  if v != "text" && v != "json" {
    return fmt.Errorf("The format \"%v\" is not allowed. Only \"text\" and \"json\" formats are allowed.", v)
  }

  return nil
}

func validateOutput(ctx *cli.Context, v string) error {
  if v != "stdout" && v != "file" {
    return fmt.Errorf("The output \"%v\" is not allowed. Only \"stdout\" and \"file\" formats are allowed.", v)
  }

  return nil
}

func validateInterval(ctx *cli.Context, v int) error {
  if !(v >= 1 && v <= 10) {
    return fmt.Errorf("The time interval must be between 1 and 10 seconds. Currently set to %v", v)
  }

  return nil
}

func generateLogs(name string, format string, output string, interval int) error {
  logger := slog.New(getHandler(format, getWriter(output)))
  num := 0
  args := []any{
    slog.String("name", name), 
    slog.String("version", "1.0.0"),
    slog.String("business_log_1", "business_log_1"),
    slog.String("business_log_2", "business_log_2"),
    slog.String("technical_log_1", "technical_log_1"),
    slog.String("technical_log_2", "technical_log_2"),
  }
  sleepTime := time.Duration(interval) * time.Second

  for true {
    num++

    switch level := 1 + rand.IntN(21); {
    case level < 3:
      logger.Error(getMessage(num), args...)
    case level < 8:
      logger.Warn(getMessage(num), args...)
    default:
      logger.Info(getMessage(num), args...)
  }

    time.Sleep(sleepTime)
  }

  return nil
}

func getWriter(output string) (io.Writer) {
  if output == "file" {
    file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

    if err != nil {
      panic(err)
    }

    return file
  }

  return os.Stdout
}

func getHandler(format string, writer io.Writer) slog.Handler {
  if format == "json" {
    return slog.NewJSONHandler(writer, nil)
  }

  return slog.NewTextHandler(writer, nil)
}

func getMessage(num int) string {
  return fmt.Sprintf("#%d log generated", num)
}
