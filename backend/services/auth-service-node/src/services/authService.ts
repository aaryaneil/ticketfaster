import bcrypt from 'bcryptjs';
import jwt from 'jsonwebtoken';
import { User } from '../models/user';
import { Credentials } from '../models/credentials';

// In-memory array to act as a user database
const users: User[] = [];

export const registerUser = async (credentials: Credentials): Promise<User> => {
  const { email, password } = credentials;

  const existingUser = users.find(user => user.email === email);
  if (existingUser) {
    throw new Error('User already exists');
  }

  const passwordHash = await bcrypt.hash(password, 10);
  
  const newUser: User = {
    id: String(users.length + 1),
    email,
    passwordHash,
  };

  users.push(newUser);
  console.log('Users in DB:', users);
  return newUser;
};

export const loginUser = async (credentials: Credentials): Promise<string | null> => {
  const { email, password } = credentials;

  const user = users.find(user => user.email === email);
  if (!user) {
    return null; // User not found
  }

  const isPasswordValid = await bcrypt.compare(password, user.passwordHash);
  if (!isPasswordValid) {
    return null; // Invalid password
  }

  const token = jwt.sign(
    { id: user.id, email: user.email },
    process.env.JWT_SECRET!,
    { expiresIn: '1h' }
  );

  return token;
};