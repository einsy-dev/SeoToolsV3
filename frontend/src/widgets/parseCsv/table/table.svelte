<script lang="ts">
  let {
    data = [
      ["Header 1", "Header 2"],
      ["Body 1", "Body 1"],
      ["Body 2", "Body 2"]
    ] as string[][]
  } = $props<{ data?: string[][] }>();
</script>

<table class="">
  <thead>
    <tr>
      {#each data[0] as el}
        <th class="border border-red-500 px-1 overflow-hidden text-nowrap">{el}</th>
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each data.slice(1, data.length) as row}
      <tr class="overflow-hidden">
        {#each row as el}
          <td class="border border-red-500 px-1 text-nowrap">{el}</td>
        {/each}
      </tr>
    {/each}
  </tbody>
</table>
