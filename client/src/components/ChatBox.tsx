import { Box, } from '@chakra-ui/react';
import React, { memo } from 'react'
import { ChatMessage } from '../types/chat.ts';

interface ChatBoxProps {
    messageList: ChatMessage[];
    currentUser: string;
}


export const ChatBox = memo((props: ChatBoxProps) => {

    const { messageList, currentUser } = props

    return (
        <Box border="1px solid gray" p="5" h="xs" mb="4" borderRadius="lg" overflowY="scroll">
            <Box>
                {
                    messageList.length > 0 && messageList.map(
                        msgItem => {

                            return (
                                <Box>
                                    <Box as="span" color={msgItem.user === currentUser ? "red" : "blue"}>{msgItem.user}</Box>: {msgItem.message}
                                </Box>
                            )
                        })
                }
            </Box>
        </Box>
    )
}, (prevProps, nextProps) => {
    if (prevProps.messageList.length === nextProps.messageList.length) {
        return true;
    }
    return false;
}) 