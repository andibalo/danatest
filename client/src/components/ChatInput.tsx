import { Button, Flex, Input } from '@chakra-ui/react';
import React from 'react'


interface ChatInputProps {
    placeholder: string;
    message: string;
    onChange: (event: any) => void;
    onSubmit: () => void;
    submitBtnText: string;
    isLoading?: boolean;
}


export const ChatInput = (props: ChatInputProps) => {

    const { onChange, onSubmit, message, placeholder, submitBtnText, isLoading } = props

    return (
        <Flex>
            <Input placeholder={placeholder} value={message} onChange={e => onChange(e)} mr="3" />
            <Button isLoading={isLoading} colorScheme='blue' onClick={onSubmit}>{submitBtnText}</Button>
        </Flex>
    )
}