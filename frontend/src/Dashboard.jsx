import { useState, useEffect } from "react";

// eslint-disable-next-line react/prop-types
export default function Dashboard({ token }) {
  // State untuk menyimpan daftar item dan input form
  const [items, setItems] = useState([]);
  const [newItemName, setNewItemName] = useState("");
  const [newItemDescription, setNewItemDescription] = useState("");

  // State untuk mengelola mode edit
  const [editingItemId, setEditingItemId] = useState(null);
  const [editingName, setEditingName] = useState("");
  const [editingDescription, setEditingDescription] = useState("");

  // URL dasar backend
  const BASE_URL = "http://192.168.1.12:8081/api/items";

  // Fungsi untuk mengambil data item dari backend
  const fetchItems = async () => {
    try {
      const res = await fetch(BASE_URL, {
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      if (res.ok) {
        const data = await res.json();
        console.log("Data items:", data); // Debug: periksa data di console
        setItems(data);
      } else {
        console.error("Gagal mengambil items");
      }
    } catch (err) {
      console.error("Error:", err);
    }
  };

  // Ambil items saat komponen di-mount
  useEffect(() => {
    fetchItems();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Fungsi untuk menambahkan item baru
  const handleAddItem = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(BASE_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          name: newItemName,
          description: newItemDescription,
        }),
      });
      if (res.status === 201) {
        const newItem = await res.json();
        // Tambahkan item baru ke state
        setItems((prevItems) => [...prevItems, newItem]);
        setNewItemName("");
        setNewItemDescription("");
      } else {
        const errorText = await res.text();
        alert("Gagal menambahkan item: " + errorText);
      }
    } catch (err) {
      console.error("Error saat menambahkan item:", err);
    }
  };

  // Fungsi untuk memulai mode edit pada item tertentu
  const startEditing = (item) => {
    setEditingItemId(item.ID);
    setEditingName(item.Name);
    setEditingDescription(item.Description);
  };

  // Fungsi untuk membatalkan mode edit
  const cancelEditing = () => {
    setEditingItemId(null);
    setEditingName("");
    setEditingDescription("");
  };

  // Fungsi untuk mengirim update item ke backend
  const handleUpdateItem = async (id) => {
    try {
      const res = await fetch(`${BASE_URL}/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          name: editingName,
          description: editingDescription,
        }),
      });
      if (res.ok) {
        const updatedItem = await res.json();
        setItems((prevItems) =>
          prevItems.map((item) => (item.ID === id ? updatedItem : item))
        );
        cancelEditing();
      } else {
        const errorText = await res.text();
        alert("Gagal mengupdate item: " + errorText);
      }
    } catch (err) {
      console.error("Error updating item:", err);
    }
  };

  // Fungsi untuk menghapus item dari backend
  const handleDeleteItem = async (id) => {
    try {
      const res = await fetch(`${BASE_URL}/${id}`, {
        method: "DELETE",
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      if (res.status === 204) {
        // Hapus item dari state
        setItems((prevItems) => prevItems.filter((item) => item.ID !== id));
      } else {
        const errorText = await res.text();
        alert("Gagal menghapus item: " + errorText);
      }
    } catch (err) {
      console.error("Error deleting item:", err);
    }
  };

  return (
    <div style={styles.container}>
      <h2>Dashboard</h2>
      <h3>Items</h3>
      <ul style={styles.itemList}>
        {items.map((item, index) => (
          <li key={item.ID ? item.ID : index} style={styles.item}>
            {editingItemId === item.ID ? (
              // Jika item sedang diedit, tampilkan form edit
              <>
                <input
                  type="text"
                  value={editingName}
                  onChange={(e) => setEditingName(e.target.value)}
                  style={styles.input}
                />
                <input
                  type="text"
                  value={editingDescription}
                  onChange={(e) => setEditingDescription(e.target.value)}
                  style={styles.input}
                />
                <button
                  onClick={() => handleUpdateItem(item.ID)}
                  style={styles.button}
                >
                  Update
                </button>
                <button onClick={cancelEditing} style={styles.cancelButton}>
                  Cancel
                </button>
              </>
            ) : (
              // Tampilkan data item
              <>
                <span style={styles.itemText}>
                  <strong>{item.Name}</strong>: {item.Description}
                </span>
                <div>
                  <button
                    onClick={() => startEditing(item)}
                    style={styles.editButton}
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDeleteItem(item.ID)}
                    style={styles.deleteButton}
                  >
                    Delete
                  </button>
                </div>
              </>
            )}
          </li>
        ))}
      </ul>
      <h3>Add New Item</h3>
      <form onSubmit={handleAddItem} style={styles.form}>
        <div style={styles.formGroup}>
          <label style={styles.label}>
            Name:
            <input
              type="text"
              value={newItemName}
              onChange={(e) => setNewItemName(e.target.value)}
              required
              style={styles.input}
            />
          </label>
        </div>
        <div style={styles.formGroup}>
          <label style={styles.label}>
            Description:
            <input
              type="text"
              value={newItemDescription}
              onChange={(e) => setNewItemDescription(e.target.value)}
              required
              style={styles.input}
            />
          </label>
        </div>
        <button type="submit" style={styles.button}>
          Add Item
        </button>
      </form>
    </div>
  );
}

// Styling menggunakan object CSS-in-JS
const styles = {
  container: {
    fontFamily: "Arial, sans-serif",
    padding: "20px",
    backgroundColor: "#f7f7f7",
    maxWidth: "800px",
    margin: "20px auto",
    borderRadius: "8px",
    boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
  },
  itemList: {
    listStyleType: "none",
    padding: 0,
    margin: "20px 0",
  },
  item: {
    backgroundColor: "#fff",
    marginBottom: "10px",
    padding: "10px",
    borderRadius: "4px",
    display: "flex",
    flexDirection: "column",
    gap: "10px",
    boxShadow: "0 1px 3px rgba(0,0,0,0.1)",
  },
  itemText: {
    flex: "1",
  },
  form: {
    backgroundColor: "#fff",
    padding: "15px",
    borderRadius: "4px",
    boxShadow: "0 1px 3px rgba(0,0,0,0.1)",
    marginTop: "20px",
  },
  formGroup: {
    marginBottom: "10px",
  },
  label: {
    display: "block",
    marginBottom: "5px",
    fontWeight: "bold",
  },
  input: {
    width: "100%",
    padding: "8px",
    borderRadius: "4px",
    border: "1px solid #ccc",
    marginBottom: "5px",
  },
  button: {
    padding: "8px 15px",
    border: "none",
    borderRadius: "4px",
    backgroundColor: "#007bff",
    color: "#fff",
    cursor: "pointer",
    marginRight: "5px",
  },
  cancelButton: {
    padding: "8px 15px",
    border: "none",
    borderRadius: "4px",
    backgroundColor: "#6c757d",
    color: "#fff",
    cursor: "pointer",
  },
  editButton: {
    padding: "6px 10px",
    border: "none",
    borderRadius: "4px",
    backgroundColor: "#ffc107",
    color: "#fff",
    cursor: "pointer",
    marginRight: "5px",
  },
  deleteButton: {
    padding: "6px 10px",
    border: "none",
    borderRadius: "4px",
    backgroundColor: "#dc3545",
    color: "#fff",
    cursor: "pointer",
  },
};
