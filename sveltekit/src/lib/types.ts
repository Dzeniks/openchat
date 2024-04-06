
export type Message = {
    SenderID: string,
    Content: string,
    SentAt: Date
}

export type Chat = {
    ChatID: string,
    OwnerID: string,
    Messages: Message[]
}