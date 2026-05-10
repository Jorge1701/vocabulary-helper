export default function FaviconImg(props: {
  url: string
}) {
  const domain = new URL(props.url).hostname;

  return (
    <img
      src={`https://www.google.com/s2/favicons?domain=${domain}&sz=32`}
      width={20}
      height={20}
      style={{ borderRadius: 4 }}
    />
  );
}
