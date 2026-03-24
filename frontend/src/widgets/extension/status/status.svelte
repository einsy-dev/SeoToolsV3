<script lang="ts">
  const {
    text = "url",
    loading = false,
    position = "1/3"
  }: { text: string; loading: boolean; position: string } = $props();
</script>

<div class="sticky top-0 bg-black rounded flex gap-2 items-center justify-center">
  <Spinner active={loading} />
  <div class="">{text}</div>
  <div class="">{position}</div>
</div>
