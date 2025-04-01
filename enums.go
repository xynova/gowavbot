package gogwave

// GGWaveSampleFormatType Data format of the audio samples
type GGWaveSampleFormatType int32

const (
	GGWaveSampleFormatUndefined GGWaveSampleFormatType = iota
	GGWaveSampleFormatU8
	GGWaveSampleFormatI8
	GGWaveSampleFormatU16
	GGWaveSampleFormatI16
	GGWaveSampleFormatF32
)

// GGWaveProtocolType Protocol ids
type GGWaveProtocolType int32

const (
	ProtocolAudibleNormal GGWaveProtocolType = iota
	ProtocolAudibleFast
	ProtocolAudibleFastest
	ProtocolUltrasoundNormal
	ProtocolUltrasoundFast
	ProtocolUltrasoundFastest
	ProtocolDTNormal
	ProtocolDTFast
	ProtocolDTFastest
	ProtocolMTNormal
	ProtocolMTFast
	ProtocolMTFastest
)

type GGWaveFilterType int32

const (
	GGWaveFilterHANN GGWaveFilterType = iota
	GGWaveFilterHamming
	GGWaveFilterFirstOrderHighPass
)

// GGWaveOperatingModeType Operating modes of ggwave
type GGWaveOperatingModeType int32

const (
	// GGWaveOperatingModeRX The instance will be able to receive audio data
	GGWaveOperatingModeRX GGWaveOperatingModeType = 1 << 1

	// GGWaveOperatingModeTX The instance will be able generate audio waveforms for transmission
	GGWaveOperatingModeTX GGWaveOperatingModeType = 1 << 2

	// GGWaveOperatingModeRXTX The instance will be able generate audio waveforms for transmission
	GGWaveOperatingModeRXTX GGWaveOperatingModeType = (GGWaveOperatingModeRX | GGWaveOperatingModeTX)

	// GGWaveOperatingModeTXOnlyTones The encoding process generates only a list of tones instead of full audio
	// waveform. This is useful for low-memory devices and embedded systems.
	GGWaveOperatingModeTXOnlyTones GGWaveOperatingModeType = 1 << 3

	// GGWaveOperatingModeUseDSS Enable the built-in Direct Sequence Spread (DSS) algorithm
	GGWaveOperatingModeUseDSS GGWaveOperatingModeType = 1 << 4
)
