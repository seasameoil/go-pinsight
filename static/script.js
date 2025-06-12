document.getElementById('bookmarkForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const url = document.getElementById('url').value;
    const note = document.getElementById('note').value;
    const tags = document.getElementById('tags').value.split(',').map(tag => tag.trim());

    fetch('/bookmarks/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url, note, tags }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('네트워크 응답이 좋지 않습니다.');
        }
        return response.json();
    })
    .then(data => {
        addBookmarkToList(data);
        document.getElementById('bookmarkForm').reset();
    })
    .catch(error => {
        console.error('Error:', error);
        alert('북마크 추가 중 오류가 발생했습니다.'); // 사용자에게 알림
    });
});

function addBookmarkToList(bookmark) {
    const list = document.getElementById('bookmarkList');
    const listItem = document.createElement('li');
    listItem.textContent = `${bookmark.url} - ${bookmark.note} [${bookmark.tags.join(', ')}]`;
    list.appendChild(listItem);
}

// 페이지 로드 시 북마크 목록 가져오기
window.onload = function() {
    fetch('/bookmarks')
        .then(response => {
            if (!response.ok) {
                throw new Error('네트워크 응답이 좋지 않습니다.');
            }
            return response.json();
        })
        .then(data => {
            data.forEach(bookmark => addBookmarkToList(bookmark));
        })
        .catch(error => {
            console.error('Error:', error);
            alert('북마크 목록을 가져오는 중 오류가 발생했습니다.'); // 사용자에게 알림
        });
};
