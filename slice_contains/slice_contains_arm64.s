#include "textflag.h"

//func SliceContainsV1(s []uint8, target uint8) bool
TEXT ·SliceContainsV1(SB), NOSPLIT, $0
    LDP slice_base+0(FP), (R0, R1) // R0 - data pointer, R1 - length
    MOVB target+24(FP), R2 // R2 = target
    VDUP R2, V1.B16 // V1 = [R2, R2, R2....R2]

loop:
    CBZ R1, no // if R1 == 0 {no code}

    VLD1.P 16(R0), [V2.B16] // V2 = *(R0)[:16]; R0 += 16
    VCMEQ V1.B16, V2.B16, V3.B16 // compare V1 and V2. V3 = [0000000 111111 000000...]

    VADDV V3.B16, V2 // V2 = sum(V3)
    VMOV V2.H[0], R4 // V2=[int16,int16,int16...]; R4 = V2[0]

    CBNZ R4, yes

    SUB $16, R1

    B loop

no:
    MOVD $0, R5
    MOVD R5, ret+32(FP)
    RET


yes:
    MOVD $1, R5
    MOVD R5, ret+32(FP)
    RET
