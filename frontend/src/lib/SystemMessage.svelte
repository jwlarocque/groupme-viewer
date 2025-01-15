<script lang="ts">
    interface Props {
        message: any;
    }

    let {
        message,
    }: Props = $props();
</script>

<style>
    .message-header {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        gap: 1em;
        font-size: 0.8em;
        color: gray;
    }
</style>

<div class="message-header">
    <div class="message-sender">
        <span>{message.name}</span>
    </div>
    <div>{message.text}</div>
    <div class="timestamp">
        {new Date(message.created_at * 1000).toLocaleString("en-US", {hour: "numeric", minute: "2-digit", hour12: false})}
    </div>
</div>