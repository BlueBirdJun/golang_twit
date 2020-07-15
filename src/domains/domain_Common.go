package domains

type ResultModel struct {
	Success  bool
	HasAlert bool
	HasError bool
	Message  string
}
