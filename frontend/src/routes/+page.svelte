<script lang="ts">
	import MessageList from "$lib/MessageList.svelte";

    async function getChats(): Promise<any> {
        const response = await fetch('/api/chats/');
        const data = await response.json();
        console.log(data);
        return data;
    }

    async function getChat(id: number): Promise<any> {
        const response = await fetch('/api/chat/' + id);
        const data = await response.json();
        console.log(data);
        return data;
    }

    function getImageUrlConverter(chatId: number): (messageId: string, urlString: string) => string {
        return (messageId: string, urlString: string): string => {
            const url = new URL(urlString);
            const imageNameParts = url.pathname.substring(1).split(".");
            const resolution = imageNameParts[0];
            const extension = imageNameParts[1];
            const id = imageNameParts[2];
            return "/data/" + chatId + "/gallery/" + messageId + "_" + resolution + "." + id + "." + extension;
        }
    }

    function getVideoUrlConverter(chatId: number): (messageId: string, urlString: string) => string {
        return (messageId: string, urlString: string): string => {
            const videoNameParts = urlString.split("/").pop()?.split(".");
            if (!videoNameParts) return "";
            const id = videoNameParts[0];
            const resolution = videoNameParts[1];
            const extension = videoNameParts[2];
            return "/data/" + chatId + "/gallery/" + messageId + "_" + id + "." + resolution + "." + extension;
        }
    }

    let chats: Promise<any> = getChats();
    let currentChatId: number | undefined = $state();
    let currentChat: any | undefined = $derived(currentChatId ? getChat(currentChatId) : undefined);
    $inspect(currentChat);
</script>

<style>
    :global(body, html) {
        position: relative;
        width: 100vw;
        height: 100vh;
        background-color: black;
        color: white;
        margin: 0;
        padding: 0;
    }

    .messages {
        max-width: 100%;
        width: 60em;
        margin: 0 auto;
    }
</style>

{#await chats}
    <p>Loading...</p>
{:then chatsData}
    <select bind:value={currentChatId}>
        <option value={undefined}>Select a chat</option>
        {#each Object.keys(chatsData) as id}
            <option value={id}>{chatsData[id]?.Metadata?.name}</option>
        {/each}
    </select>
    {#if currentChatId}
        {#await currentChat}
            <span>Loading...</span>
            <span>(this may take a while)</span>
        {:then chatData}
            <span> - {chatData.Metadata.description}</span>
            <div class="messages">
                <MessageList
                    messages={chatData.Messages}
                    convertImageUrl={getImageUrlConverter(currentChatId)}
                    convertVideoUrl={getVideoUrlConverter(currentChatId)}
                    users={chatData.Metadata.members.reduce((dict: any, user: any) => { 
                        dict[user.user_id] = user;
                        return dict;
                    }, {})}
                />
            </div>
        {/await}
    {/if}
{/await}