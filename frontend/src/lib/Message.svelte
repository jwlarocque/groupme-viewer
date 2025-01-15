<script lang="ts">
	import ImageAttachment from "./ImageAttachment.svelte";
	import VideoAttachment from "./VideoAttachment.svelte";

    const GROUP_TIME_DELTA = 5 * 60;

    interface Props {
        users: any;
        message: any;
        lastMessage: any;
        convertImageUrl: (messageId: string, urlString: string) => string;
        convertVideoUrl: (messageId: string, urlString: string) => string;
    }

    let {
        users,
        message,
        lastMessage,
        convertImageUrl,
        convertVideoUrl,
    }: Props = $props();
</script>

<style>
    .message-header {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        gap: 1em;
    }

    .message-content {
        display: flex;
        flex-direction: column;
        gap: 1em;
        margin: 0;
        background-color: rgb(26, 24, 22);
        border-radius: 0.5em;
        padding: 1em;
    }

    .message-content > p {
        margin: 0;
    }

    .attachments {
        display: flex;
        flex-direction: row;
        gap: 1em;
    }

    .timestamp {
        font-size: 0.8em;
        color: gray;
    }

    .reaction {
        display: flex;
        flex-direction: row;
        line-height: 1;
        cursor: default;
        width: fit-content;
    }

    .reactors {
        display: none;
        margin-left: 0.5em;
        cursor: auto;
    }

    .reaction:hover .reactors {
        display: block;
    }
</style>

<div class="message" id={message.id}>
    {#if !lastMessage || message.created_at - lastMessage.created_at > GROUP_TIME_DELTA || message.sender_id !== lastMessage.sender_id}
        <div class="message-header">
            <div class="message-sender">
                <span>{message.name}{users[message.sender_id]?.name != message.name ? ` (${users[message.sender_id]?.name || "inactive"})` : ""}</span>
            </div>
            <div class="timestamp">
                {new Date(message.created_at * 1000).toLocaleString("en-US", {hour: "numeric", minute: "2-digit", hour12: false})}
            </div>
        </div>
    {/if}
    <div class="message-content">
        {#if message?.attachments?.filter((attachment: any) => attachment.type !== "mentions").length > 0}
            <div class="attachments">
                {#each message.attachments as attachment}
                    {#if attachment.type === "image"}
                        <ImageAttachment url={convertImageUrl(message.id, attachment.url)} />
                    {:else if attachment.type === "mentions"}
                    {:else if attachment.type === "video"}
                        <VideoAttachment url={convertVideoUrl(message.id, attachment.url)} />
                    {:else if attachment.type === "reply"}
                        <p>Reply to: <a href="#{attachment.reply_id}">{users[(attachment.user_id).toString()]?.name} - {attachment.reply_id}</a></p> 
                    {:else}
                        <p>{attachment.type}</p>
                        <p>{attachment.url}</p>
                    {/if}
                {/each}
            </div>
        {/if}
        {#if message.text}
            <p>{message.text}</p>
        {/if}
        {#if message?.reactions?.length > 0}
            <div class="reactions">
                {#each message.reactions as reaction}
                    <div class="reaction" title={reaction.user_ids.map((userId: string) => users[userId]?.name).join(", ")}>
                        {reaction?.code || "❤️"} {reaction.user_ids.length}
                        <span class="reactors"> - {reaction.user_ids.map((userId: string) => users[userId]?.name).join(", ")}</span>
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</div>