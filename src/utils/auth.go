package utils


func ValidateLoginSession(dbSession, tokenSession string) bool {
    return dbSession == tokenSession
}
