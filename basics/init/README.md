# `init()` Function in Go

The `init()` function is a **special function** in Go that is automatically invoked by the Go runtime:

- It runs **before the `main()` function**
- It is typically used to **initialize state**, **set up configuration**, **register components**, or **validate environment setup**

---

## ✅ Key Properties of `init()` Function

| Feature                         | Description                                    |
|----------------------------     |------------------------------------------------|
| **Signature**                   | `func init()` – no parameters, no return values |
| **Auto-invoked**                | You never call it yourself                     |
| **One per file allowed**        | Each Go file in a package can have its own `init()` |
| **Runs before `main()`**        | Called after all `var` initializations         |
| **Can be multiple per package** | They execute in the order they appear across files |
| **Cannot be exported**          | Not accessible from outside the package        |

---

## 🔧 Common Use Cases

- Initialize global variables
- Read configuration or environment variables
- Register plugins (like in image, database drivers)
- Connect to database
- Set up logging

---

## 🚨 Important Notes

- You **shouldn't abuse `init()`** — keep it simple and predictable.
- Avoid logic that’s **hard to test or debug**.
- Prefer **explicit initialization** via constructor functions when possible.

---

## 📦 Multiple `init()` in a Package

If you have several files in a package with their own `init()` functions, they will all run — in the order the Go compiler determines (usually based on file dependency order).

---

## ✅ Summary

| Point            | Description                                     |
|------------------|-------------------------------------------------|
| **Purpose**      | Auto setup before `main()`                      |
| **Signature**    | `func init()` with no args or returns           |
| **When runs**    | After global `var` init, before `main()`        |
| **Common Uses**  | Config loading, setup, registration             |
