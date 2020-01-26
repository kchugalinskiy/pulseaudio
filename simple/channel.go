package simple

// #include <pulse/simple.h>
// #include <pulse/channelmap.h>
import "C"
import (
	"fmt"
)

// ChannelPosition is a basic channel position enum. It is being actively used in mapping between channel index and its
// physical destination (e.g. front left speaker)
// See https://www.freedesktop.org/software/pulseaudio/doxygen/channelmap_8h.html#af1cbe2738487c74f99e613779bd34bf2
type ChannelPosition int

const (
	ChannelPositionInvalid ChannelPosition = iota
	ChannelPositionMono
	ChannelPositionFrontLeft
	ChannelPositionFrontRight
	ChannelPositionFrontCenter
	ChannelPositionRearCenter
	ChannelPositionRearLeft
	ChannelPositionRearRight
	ChannelPositionLFE
	ChannelPositionFrontLeftOfCenter
	ChannelPositionFrontRightOfCenter
	ChannelPositionSideLeft
	ChannelPositionSideRight
	ChannelPositionAUX0
	ChannelPositionAUX1
	ChannelPositionAUX2
	ChannelPositionAUX3
	ChannelPositionAUX4
	ChannelPositionAUX5
	ChannelPositionAUX6
	ChannelPositionAUX7
	ChannelPositionAUX8
	ChannelPositionAUX9
	ChannelPositionAUX10
	ChannelPositionAUX11
	ChannelPositionAUX12
	ChannelPositionAUX13
	ChannelPositionAUX14
	ChannelPositionAUX15
	ChannelPositionAUX16
	ChannelPositionAUX17
	ChannelPositionAUX18
	ChannelPositionAUX19
	ChannelPositionAUX20
	ChannelPositionAUX21
	ChannelPositionAUX22
	ChannelPositionAUX23
	ChannelPositionAUX24
	ChannelPositionAUX25
	ChannelPositionAUX26
	ChannelPositionAUX27
	ChannelPositionAUX28
	ChannelPositionAUX29
	ChannelPositionAUX30
	ChannelPositionAUX31
	ChannelPositionTopCenter
	ChannelPositionTopFrontLeft
	ChannelPositionTopFrontRight
	ChannelPositionTopFrontCenter
	ChannelPositionTopRearLeft
	ChannelPositionTopRearRight
	ChannelPositionTopRearCenter
	ChannelPositionMax
)

func toChannel(p ChannelPosition) C.pa_channel_position_t {
	switch p {
	case ChannelPositionInvalid:
		return C.PA_CHANNEL_POSITION_INVALID
	case ChannelPositionMono:
		return C.PA_CHANNEL_POSITION_MONO
	case ChannelPositionFrontLeft:
		return C.PA_CHANNEL_POSITION_FRONT_LEFT
	case ChannelPositionFrontRight:
		return C.PA_CHANNEL_POSITION_FRONT_RIGHT
	case ChannelPositionFrontCenter:
		return C.PA_CHANNEL_POSITION_FRONT_CENTER
	case ChannelPositionRearCenter:
		return C.PA_CHANNEL_POSITION_REAR_CENTER
	case ChannelPositionRearLeft:
		return C.PA_CHANNEL_POSITION_REAR_LEFT
	case ChannelPositionRearRight:
		return C.PA_CHANNEL_POSITION_REAR_RIGHT
	case ChannelPositionLFE:
		return C.PA_CHANNEL_POSITION_LFE
	case ChannelPositionFrontLeftOfCenter:
		return C.PA_CHANNEL_POSITION_FRONT_LEFT_OF_CENTER
	case ChannelPositionFrontRightOfCenter:
		return C.PA_CHANNEL_POSITION_FRONT_RIGHT_OF_CENTER
	case ChannelPositionSideLeft:
		return C.PA_CHANNEL_POSITION_SIDE_LEFT
	case ChannelPositionSideRight:
		return C.PA_CHANNEL_POSITION_SIDE_RIGHT
	case ChannelPositionAUX0:
		return C.PA_CHANNEL_POSITION_AUX0
	case ChannelPositionAUX1:
		return C.PA_CHANNEL_POSITION_AUX1
	case ChannelPositionAUX2:
		return C.PA_CHANNEL_POSITION_AUX2
	case ChannelPositionAUX3:
		return C.PA_CHANNEL_POSITION_AUX3
	case ChannelPositionAUX4:
		return C.PA_CHANNEL_POSITION_AUX4
	case ChannelPositionAUX5:
		return C.PA_CHANNEL_POSITION_AUX5
	case ChannelPositionAUX6:
		return C.PA_CHANNEL_POSITION_AUX6
	case ChannelPositionAUX7:
		return C.PA_CHANNEL_POSITION_AUX7
	case ChannelPositionAUX8:
		return C.PA_CHANNEL_POSITION_AUX8
	case ChannelPositionAUX9:
		return C.PA_CHANNEL_POSITION_AUX9
	case ChannelPositionAUX10:
		return C.PA_CHANNEL_POSITION_AUX10
	case ChannelPositionAUX11:
		return C.PA_CHANNEL_POSITION_AUX11
	case ChannelPositionAUX12:
		return C.PA_CHANNEL_POSITION_AUX12
	case ChannelPositionAUX13:
		return C.PA_CHANNEL_POSITION_AUX13
	case ChannelPositionAUX14:
		return C.PA_CHANNEL_POSITION_AUX14
	case ChannelPositionAUX15:
		return C.PA_CHANNEL_POSITION_AUX15
	case ChannelPositionAUX16:
		return C.PA_CHANNEL_POSITION_AUX16
	case ChannelPositionAUX17:
		return C.PA_CHANNEL_POSITION_AUX17
	case ChannelPositionAUX18:
		return C.PA_CHANNEL_POSITION_AUX18
	case ChannelPositionAUX19:
		return C.PA_CHANNEL_POSITION_AUX19
	case ChannelPositionAUX20:
		return C.PA_CHANNEL_POSITION_AUX20
	case ChannelPositionAUX21:
		return C.PA_CHANNEL_POSITION_AUX21
	case ChannelPositionAUX22:
		return C.PA_CHANNEL_POSITION_AUX22
	case ChannelPositionAUX23:
		return C.PA_CHANNEL_POSITION_AUX23
	case ChannelPositionAUX24:
		return C.PA_CHANNEL_POSITION_AUX24
	case ChannelPositionAUX25:
		return C.PA_CHANNEL_POSITION_AUX25
	case ChannelPositionAUX26:
		return C.PA_CHANNEL_POSITION_AUX26
	case ChannelPositionAUX27:
		return C.PA_CHANNEL_POSITION_AUX27
	case ChannelPositionAUX28:
		return C.PA_CHANNEL_POSITION_AUX28
	case ChannelPositionAUX29:
		return C.PA_CHANNEL_POSITION_AUX29
	case ChannelPositionAUX30:
		return C.PA_CHANNEL_POSITION_AUX30
	case ChannelPositionAUX31:
		return C.PA_CHANNEL_POSITION_AUX31
	case ChannelPositionTopCenter:
		return C.PA_CHANNEL_POSITION_TOP_CENTER
	case ChannelPositionTopFrontLeft:
		return C.PA_CHANNEL_POSITION_TOP_FRONT_LEFT
	case ChannelPositionTopFrontRight:
		return C.PA_CHANNEL_POSITION_TOP_FRONT_RIGHT
	case ChannelPositionTopFrontCenter:
		return C.PA_CHANNEL_POSITION_TOP_FRONT_CENTER
	case ChannelPositionTopRearLeft:
		return C.PA_CHANNEL_POSITION_TOP_REAR_LEFT
	case ChannelPositionTopRearRight:
		return C.PA_CHANNEL_POSITION_TOP_REAR_RIGHT
	case ChannelPositionTopRearCenter:
		return C.PA_CHANNEL_POSITION_TOP_REAR_CENTER
	case ChannelPositionMax:
		return C.PA_CHANNEL_POSITION_MAX
	}
	panic(fmt.Errorf("unsupported channel position value: %v", p))
}
