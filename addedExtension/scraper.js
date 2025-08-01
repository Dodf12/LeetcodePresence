
function scrapeData() {
  console.log("[scraper] Starting scrapeData...");
  const title = document.title;
  console.log("[scraper] Got title:", title);
  const slug = window.location.pathname.split("/")[2] || "unknown";
  console.log("[scraper] Got slug:", slug);
  const difficulty = document.querySelector('[class*="text-difficulty-"]')?.innerText || "Unknown";
  console.log("[scraper] Got difficulty:", difficulty);
  const lang = document.querySelector(".ant-select-selection-item")?.innerText || "Unknown";
  console.log("[scraper] Got language:", lang);
  const status = document.querySelector(".status__3q2d")?.innerText || "Not Submitted";
  console.log("[scraper] Got status:", status);

  // Extract problem number from title (e.g., "123. Two Sum - LeetCode")
  let problemNumber = "";
  const match = title.match(/^(\d+)\./);
  if (match) {
    problemNumber = match[1];
  } else {
    problemNumber = "";
  }
  console.log("[scraper] Got problemNumber:", problemNumber);

  const data = { title, slug, difficulty, language: lang, status, problemNumber };
  console.log("[scraper] Prepared data:", data);

  // Send to your local Go server
  const url = "http://localhost:8080/posts";
  console.log(`[scraper] Sending POST to ${url} with data:`, data);
  console.log("[scraper] About to call fetch...");
  fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  })
    .then(response => {
      console.log("[scraper] Fetch response status:", response.status);
      return response.text();
    })
    .then(text => {
      console.log("[scraper] Fetch response body:", text);
      console.log("[scraper] Fetch call finished (then)");
    })
    .catch(err => {
      console.error("[scraper] Fetch error:", err);
      console.log("[scraper] Fetch call finished (catch)");
    });
  console.log("[scraper] Fetch call initiated");
}

// ✅ Scrape every 15 seconds
setInterval(scrapeData, 15000);

// Or scrape once right away too
scrapeData();
