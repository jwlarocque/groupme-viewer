<script lang="ts">
    interface Props {
        url: string;
    }

    let {
        url,
    }: Props = $props();

    let overlay = $state(false);
</script>

<!-- this is an horrific hack -->
<svelte:window onkeydown={(e) => {if (overlay && e.key === 'Escape') overlay = false}} />

<style>
    img {
        max-width: 10em;
        max-height: 10em;
        object-fit: contain;
        border-radius: 0.5em;
        cursor: pointer;
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

<img src={url} loading="lazy" onclick={() => (overlay = !overlay)} />
{#if overlay}
    <div class="overlay"
        onclick={() => (overlay = !overlay)}
    >
        <img src={url} loading="lazy" onclick={(e) => e.stopPropagation()} />
    </div>
{/if}