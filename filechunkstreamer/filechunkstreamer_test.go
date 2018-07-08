package filechunkstreamer

import "testing"

func TestSample(t *testing.T) {
}

func BenchmarkSmallChunkStreamerNarrow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StreamFileChunks(25,1,"/tmp/dat")
	}
}
func benchmarkSmallChunkStreamerWide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StreamFileChunks(1024,128,"/tmp/dat")
	}
}

func benchmarkLargeChunkStreamerNarrow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StreamFileChunks(25,1,"/home/jeremiah/Downloads/avatar.2009.extended.multi.1080p.mkv")
	}
}
func benchmarkLargeChunkStreamerWide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StreamFileChunks(1024,128,"/home/jeremiah/Downloads/avatar.2009.extended.multi.1080p.mkv")
	}
}

func benchmarkLargeChunkStreamerUltraWide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StreamFileChunks(2048,1536,"/home/jeremiah/Downloads/avatar.2009.extended.multi.1080p.mkv")
	}
}
