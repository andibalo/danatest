import { client } from "./client.ts"

const USER_BASE_PATH = "/user"

export const registerUser = async (username : string) => {
    return client.post(USER_BASE_PATH + "/", { username })
}

export const deleteUser = async (username : string) => {
    return client.delete(USER_BASE_PATH + "/", { username })
}