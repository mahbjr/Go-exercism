package account

import (
	"sync"
)

// Account representa uma conta bancária
type Account struct {
	balance int64
	closed  bool
	mu      sync.Mutex
}

// Open cria uma nova conta com o saldo inicial especificado
// Retorna nil se o saldo inicial for negativo
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance: amount,
		closed:  false,
	}
}

// Balance retorna o saldo atual da conta
// O segundo valor retornado é false se a conta estiver fechada, caso contrário é true
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit adiciona ou remove fundos da conta
// Retorna o novo saldo e true se a operação for bem-sucedida
// Retorna 0 e false se a conta estiver fechada ou se a operação resultaria em saldo negativo
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Verifica se a conta está fechada
	if a.closed {
		return 0, false
	}

	// Calcula o novo saldo
	newBalance := a.balance + amount

	// Não permite saldo negativo
	if newBalance < 0 {
		return 0, false
	}

	// Atualiza o saldo
	a.balance = newBalance

	return a.balance, true
}

// Close fecha a conta e retorna o saldo final
// O segundo valor retornado é true se a conta foi fechada com sucesso
// Retorna 0 e false se a conta já estiver fechada
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Verifica se a conta já está fechada
	if a.closed {
		return 0, false
	}

	// Armazena o saldo atual
	payout := a.balance

	// Fecha a conta e zera o saldo
	a.closed = true
	a.balance = 0

	return payout, true
}
