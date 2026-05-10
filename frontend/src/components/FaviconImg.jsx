export default function FaviconImg({ url }) {
  const domain = new URL(url).hostname;

  return (
    <img
      src={`https://www.google.com/s2/favicons?domain=${domain}&sz=32`}
      width={20}
      height={20}
      style={{ borderRadius: 4 }}
    />
  );
}
