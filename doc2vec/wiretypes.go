package doc2vec

import (
	"sync"

	"github.com/topxeq/doc2vec/corpus"
	"github.com/topxeq/doc2vec/neuralnet"
)

//go:generate msgp

type SortItem struct {
	Idx int32
	Dis float64
}

type TSortItemSlice []*SortItem

type IDoc2Vec interface {
	Train(fname string)
	GetCorpus() corpus.ICorpus
	GetNeuralNet() neuralnet.INeuralNet
	SaveModel(fname string) (err error)
	LoadModel(fname string) (err error)
	Word2Words(word string)
	TXWord2Words(word string, limit int, outWordsA map[string]int) []string
	TXWord2WordsIn(word string, limit int, inWordListA map[string]int) []string
	TXWord2WordsInOut(word string, limit int, inWordListA map[string]int, outWordsA map[string]int) []string
	TXWord2Docs(word string, limit int) []string
	Word2Docs(word string)
	Sen2Words(content string, iters int)
	InferDoc(content string, iters int) (rs []float32)
	Sen2Docs(content string, iters int)
	TXDoc2Words(content string, iters int, limit int, outWordsA map[string]int) (rs []string, vec []float32)
	FindNearestDoc(content string, iters int) string
	Doc2Docs(docidx int)
	Doc2Words(docidx int)
	GetLikelihood4Doc(context string) (likelihood float64)
	GetLeaveOneOutKwds(content string, iters int)
	DocSimCal(content1 string, content2 string) (dis float64)
}

type TDoc2VecImpl struct {
	Trainfile    string
	Dim          int
	UseCbow      bool //true:Continuous Bag-of-Word Model false:skip-gram
	WindowSize   int  //cbow model的窗口大小
	UseHS        bool
	UseNEG       bool //UseHS / UseNEG两种求解优化算法必须选一个 也可以两种算法都选 详见google word2vec源代码
	Negative     int  //负采样词的个数
	StartAlpha   float64
	Iters        int
	TrainedWords int
	Corpus       corpus.ICorpus
	NN           neuralnet.INeuralNet
	Pool         *sync.Pool
}
