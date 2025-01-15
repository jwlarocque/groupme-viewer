<script lang="ts">
	import ImageAttachment from "./ImageAttachment.svelte";
	import Message from "./Message.svelte";
	import SystemMessage from "./SystemMessage.svelte";

    const DISPLAY_TIME_DELTA = 60 * 60;

    interface Props {
        messages: any[];
        convertImageUrl: (messageId: string, urlString: string) => string;
        convertVideoUrl: (messageId: string, urlString: string) => string;
        users: any;
    }

    let {
        messages,
        convertImageUrl,
        convertVideoUrl,
        users,
    }: Props = $props();
</script>

<style>
    main {
        display: flex;
        gap: 0.5em;
        flex-direction: column;
    }

    .timestamp {
        font-size: 0.8em;
        color: gray;
        width: 100%;
        text-align: center;
    }
</style>

<main>
    {#each messages as message, i}
        {#if i === 0
            || message.created_at - messages[i - 1].created_at > DISPLAY_TIME_DELTA
            || new Date(message.created_at * 1000).getDay() !== new Date(messages[i - 1].created_at * 1000).getDay()
        }
            <div class="timestamp timestamp-standalone">
                - {new Date(message.created_at * 1000).toLocaleString("en-US", {
                    weekday: "long", year: "numeric", month: "long", day: "numeric",
                    hour: "numeric", minute: "2-digit", hour12: false})} -
            </div>
        {/if}
        {#if message.sender_id === "system"}
            <SystemMessage message={message} />
        {:else}
            <Message
                message={message}
                lastMessage={messages[i - 1] ?? undefined}
                convertImageUrl={convertImageUrl}
                convertVideoUrl={convertVideoUrl}
                users={users}
            />
        {/if}
    {/each}
</main>