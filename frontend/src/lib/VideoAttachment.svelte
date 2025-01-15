<script lang="ts">
    interface Props {
        url: string;
    }

    let {
        url,
    }: Props = $props();
</script>

<style>
    video {
        max-width: 20em;
        max-height: 20em;
        object-fit: contain;
        border-radius: 0.5em;
    }

    .overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
    }

    .overlay img {
        position: relative;
        display: block;
        box-sizing: border-box;
        height: fit-content;
        max-height: calc(100vh - 4em);
        width: auto;
        max-width: calc(100vw - 4em);
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        cursor: auto;
    }
</style>

<video src={url} controls preload="none"></video>