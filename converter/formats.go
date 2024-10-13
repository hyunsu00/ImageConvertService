package converter

var SupportedInputFormats = []string{"jpeg", "png", "gif", "tiff", "svg"}
var SupportedOutputFormats = []string{"jpeg", "png", "gif", "tiff"}

func IsSupportedInputFormat(format string) bool {
	for _, f := range SupportedInputFormats {
		if f == format {
			return true
		}
	}
	return false
}

func IsSupportedOutputFormat(format string) bool {
	for _, f := range SupportedOutputFormats {
		if f == format {
			return true
		}
	}
	return false
}
