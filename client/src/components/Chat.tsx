import React, { useState } from 'react'
import { Flex, Box } from '@chakra-ui/react';
import { ChatInput } from './ChatInput.tsx';
import { ChatBox } from './ChatBox.tsx';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { ChatMessage } from '../types/chat.ts';
import { userService } from '../services/index.ts';
import { ERR_CODE } from '../constants/errors.ts';
import { toast } from 'react-toastify';

export const Chat = () => {
    const [isUsernameSubmissionLoading, setIsUsernameSubmissionLoading] = useState(false)
    const [currentUser, setCurrentUser] = useState("")
    const [message, setMessage] = useState("")
    const [username, setUsername] = useState("")
    const [messageList, setMessageList] = useState<ChatMessage[]>([])

    const { lastMessage, readyState, sendJsonMessage } = useWebSocket('ws://127.0.0.1:8080/ws', {
        onOpen: () => console.log('opened'),
        onMessage: (message) => {
            setMessageList([...messageList, JSON.parse(message.data)])
        },
        shouldReconnect: (closeEvent) => true,
    });


    const onMessageChange = (e) => {
        setMessage(e.target.value)
    }

    const onMessageSubmit = () => {
        if (message === "") {
            return
        }

        sendJsonMessage({
            user: currentUser,
            message: message
        })

        setMessage("")
    }

    const onUsernameChange = (e) => {
        setUsername(e.target.value)
    }

    const onUsernameSubmit = async () => {
        if (username === "") {
            return
        }

        try {
            setIsUsernameSubmissionLoading(true)
            const res = await userService.registerUser(username)

            setCurrentUser(username)

            setUsername("")
            setIsUsernameSubmissionLoading(false)
        } catch (error) {
            console.log(error)

            if (error.response.data.code === ERR_CODE.DUPLICATE_USER) {
                toast.error("Username already exists")
                return
            }

            toast.error(error.message)

            setIsUsernameSubmissionLoading(false)
        }
    }


    return (
        <Box>
            <ChatBox messageList={messageList} currentUser={currentUser} />
            {
                currentUser ?
                    <ChatInput
                        submitBtnText="Send"
                        placeholder="Send message..."
                        message={message}
                        onChange={onMessageChange}
                        onSubmit={onMessageSubmit} /> :
                    <ChatInput
                        isLoading={isUsernameSubmissionLoading}
                        submitBtnText="Submit"
                        placeholder="Set your username"
                        message={username}
                        onChange={onUsernameChange}
                        onSubmit={onUsernameSubmit} />
            }
        </Box>
    )
}