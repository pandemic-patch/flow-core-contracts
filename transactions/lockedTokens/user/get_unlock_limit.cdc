import LockedTokens from 0xLOCKEDTOKENADDRESS

pub fun main(account: Address): UFix64 {

    let lockedAccountInfoRef = getAccount(account)
        .getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(LockedTokens.LockedAccountInfoPublicPath)!
        .borrow() ?? panic("Could not borrow a reference to public LockedAccountInfo")

    return lockedAccountInfoRef.getUnlockLimit()
}
