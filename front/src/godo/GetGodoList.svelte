<script lang="ts">
    import { each } from "svelte/internal";

    const myUrl = "http://127.0.0.1:4081/godoList";

    let result = [];

    async function postGodo() {
        const res = await fetch(myUrl, {
            method: "GET",
        });
        const json = await res.json();

        result = Object.values(json["godoList"]);
        console.log(result);
    }
</script>

<main>
    <legend>get godo List</legend>
    <button type="button" on:click|preventDefault={postGodo}>Get</button>

    <ul>
        {#each result as unitgodo, i}
            <li>
                <!-- {unitgodo.Title} -->
                {unitgodo["Title"]}
            </li>
        {/each}
    </ul>
</main>

<style>
</style>
